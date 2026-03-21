package nitrogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/AdamJCrawford/NitroGo/nitrogo/models"
)

const (
	appFWArchiveURL                              = "/nitro/v1/config/appfwarchive"
	appFWConfidFieldURL                          = "/nitro/v1/config/appfwconfidfield"
	appFWCustomSettingsURL                       = "/nitro/v1/config/appfwcustomsettings"
	appFWFieldTypeURL                            = "/nitro/v1/config/appfwfieldtype"
	appFWGlobalAppFWPolicyBindingURL             = "/nitro/v1/config/appfwglobal_appfwpolicy_binding"
	appFWGlobalAuditNSLogPolicyBindingURL        = "/nitro/v1/config/appfwglobal_auditnslogpolicy_binding"
	appFWGlobalAuditSyslogPolicyBindingURL       = "/nitro/v1/config/appfwglobal_auditsyslogpolicy_binding"
	appFWGlobalBindingURL                        = "/nitro/v1/config/appfwglobal_binding"
	appFWHTMLErrorPageURL                        = "/nitro/v1/config/appfwhtmlerrorpage"
	appFWJSONContentTypeURL                      = "/nitro/v1/config/appfwjsoncontenttype"
	appFWJSONErrorPageURL                        = "/nitro/v1/config/appfwjsonerrorpage"
	appFWLearningDataURL                         = "/nitro/v1/config/appfwlearningdata"
	appFWLearningSettingsURL                     = "/nitro/v1/config/appfwlearningsettings"
	appFWMultipartFormContentTypeURL             = "/nitro/v1/config/appfwmultipartformcontenttype"
	appFWPolicyURL                               = "/nitro/v1/config/appfwpolicy"
	appFWPolicyAppFWGlobalBindingURL             = "/nitro/v1/config/appfwpolicy_appfwglobal_binding"
	appFWPolicyAppFWPolicyLabelBindingURL        = "/nitro/v1/config/appfwpolicy_appfwpolicylabel_binding"
	appFWPolicyBindingURL                        = "/nitro/v1/config/appfwpolicy_binding"
	appFWPolicyCSVServerBindingURL               = "/nitro/v1/config/appfwpolicy_csvserver_binding"
	appFWPolicyLBVServerBindingURL               = "/nitro/v1/config/appfwpolicy_lbvserver_binding"
	appFWPolicyLabelURL                          = "/nitro/v1/config/appfwpolicylabel"
	appFWPolicyLabelAppFWPolicyBindingURL        = "/nitro/v1/config/appfwpolicylabel_appfwpolicy_binding"
	appFWPolicyLabelBindingURL                   = "/nitro/v1/config/appfwpolicylabel_binding"
	appFWPolicyLabelPolicyBindingBindingURL      = "/nitro/v1/config/appfwpolicylabel_policybinding_binding"
	appFWProfileURL                              = "/nitro/v1/config/appfwprofile"
	appFWProfileBindingURL                       = "/nitro/v1/config/appfwprofile_binding"
	appFWProfileCMDInjectionBindingURL           = "/nitro/v1/config/appfwprofile_cmdinjection_binding"
	appFWProfileContentTypeBindingURL            = "/nitro/v1/config/appfwprofile_contenttype_binding"
	appFWProfileCookieConsistencyBindingURL      = "/nitro/v1/config/appfwprofile_cookieconsistency_binding"
	appFWProfileCreditCardNumberBindingURL       = "/nitro/v1/config/appfwprofile_creditcardnumber_binding"
	appFWProfileCrossSiteScriptingBindingURL     = "/nitro/v1/config/appfwprofile_crosssitescripting_binding"
	appFWProfileCSRFTagBindingURL                = "/nitro/v1/config/appfwprofile_csrftag_binding"
	appFWProfileDenyURLBindingURL                = "/nitro/v1/config/appfwprofile_denyurl_binding"
	appFWProfileExcludeRESContentTypeBindingURL  = "/nitro/v1/config/appfwprofile_excluderescontenttype_binding"
	appFWProfileFieldConsistencyBindingURL       = "/nitro/v1/config/appfwprofile_fieldconsistency_binding"
	appFWProfileFieldFormatBindingURL            = "/nitro/v1/config/appfwprofile_fieldformat_binding"
	appFWProfileFileUploadTypeBindingURL         = "/nitro/v1/config/appfwprofile_fileuploadtype_binding"
	appFWProfileJSONDOSURLBindingURL             = "/nitro/v1/config/appfwprofile_jsondosurl_binding"
	appFWProfileJSONSQLURLBindingURL             = "/nitro/v1/config/appfwprofile_jsonsqlurl_binding"
	appFWProfileJSONXSSURLBindingURL             = "/nitro/v1/config/appfwprofile_jsonxssurl_binding"
	appFWProfileLogExpressionBindingURL          = "/nitro/v1/config/appfwprofile_logexpression_binding"
	appFWProfileSafeObjectBindingURL             = "/nitro/v1/config/appfwprofile_safeobject_binding"
	appFWProfileSQLInjectionBindingURL           = "/nitro/v1/config/appfwprofile_sqlinjection_binding"
	appFWProfileStartURLBindingURL               = "/nitro/v1/config/appfwprofile_starturl_binding"
	appFWProfileTrustedLearningClientsBindingURL = "/nitro/v1/config/appfwprofile_trustedlearningclients_binding"
	appFWProfileXMLAttachmentURLBindingURL       = "/nitro/v1/config/appfwprofile_xmlattachmenturl_binding"
	appFWProfileXMLDOSURLBindingURL              = "/nitro/v1/config/appfwprofile_xmldosurl_binding"
	appFWProfileXMLSQLInjectionBindingURL        = "/nitro/v1/config/appfwprofile_xmlsqlinjection_binding"
	appFWProfileXMLValidationURLBindingURL       = "/nitro/v1/config/appfwprofile_xmlvalidationurl_binding"
	appFWProfileXMLWSIURLBindingURL              = "/nitro/v1/config/appfwprofile_xmlwsiurl_binding"
	appFWProfileXMLXSSBindingURL                 = "/nitro/v1/config/appfwprofile_xmlxss_binding"
	appFWSettingsURL                             = "/nitro/v1/config/appfwsettings"
	appFWSignaturesURL                           = "/nitro/v1/config/appfwsignatures"
	appFWTransactionRecordsURL                   = "/nitro/v1/config/appfwtransactionrecords"
	appFWURLEncodedFormContentTypeURL            = "/nitro/v1/config/appfwurlencodedformcontenttype"
	appFWWSDLURL                                 = "/nitro/v1/config/appfwwsdl"
	appFWXMLContentTypeURL                       = "/nitro/v1/config/appfwxmlcontenttype"
	appFWXMLErrorPageURL                         = "/nitro/v1/config/appfwxmlerrorpage"
	appFWXMLSchemaURL                            = "/nitro/v1/config/appfwxmlschema"
)

// Application Firewall configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfw
type AppFWService struct {
	client *Client
}

// appfwarchive
// Configuration for archive resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwarchive
func (s *AppFWService) DeleteAppFWArchive(name string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWArchiveURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) ExportAppFWArchive(resource models.AppFWArchive) error {
	payload := map[string]any{
		"appfwarchive": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWArchiveURL+"?action=export", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) ImportAppFWArchive(resource models.AppFWArchive) error {
	payload := map[string]any{
		"appfwarchive": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWArchiveURL+"?action=Import", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWArchive() ([]models.AppFWArchive, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWArchiveURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Archives []models.AppFWArchive `json:"appfwarchive"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Archives, nil
}

// appfwconfidfield
// Configuration for configured confidential form fields resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwconfidfield
func (s *AppFWService) AddAppFWConfigField(resource models.AppFWConfidField) error {
	payload := map[string]any{
		"appfwconfidfield": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWConfidFieldURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWConfigField(name string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWConfidFieldURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) UpdateAppFWConfigField(resource models.AppFWConfidField) error {
	payload := map[string]any{
		"appfwconfidfield": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, appFWConfidFieldURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) UnsetAppFWConfigField(resource models.AppFWConfidField) error {
	payload := map[string]any{
		"appfwconfidfield": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWConfidFieldURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWConfigField() ([]models.AppFWConfidField, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWConfidFieldURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Fields []models.AppFWConfidField `json:"appfwconfidfield"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Fields, nil
}

func (s *AppFWService) CountAppFWConfigField() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWConfidFieldURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Fields []struct {
			Count float64 `json:"__count"`
		} `json:"appfwconfidfield"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Fields) > 0 {
		return result.Fields[0].Count, nil
	}

	return 0, nil
}

// appfwcustomsettings
// Configuration for application firewall custom settings XML configuration resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwcustomsettings
func (s *AppFWService) ExportAppFWCustomSettings(resource models.AppFWCustomSettings) error {
	payload := map[string]any{
		"appfwcustomsettings": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWCustomSettingsURL+"?action=export", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// appfwfieldtype
// Configuration for application firewall form field type resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwfieldtype
func (s *AppFWService) AddAppFWFieldType(resource models.AppFWFieldType) error {
	payload := map[string]any{
		"appfwfieldtype": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWFieldTypeURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWFieldType(name string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWFieldTypeURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) UpdateAppFWFieldType(resource models.AppFWFieldType) error {
	payload := map[string]any{
		"appfwfieldtype": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, appFWFieldTypeURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWFieldType() ([]models.AppFWFieldType, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWFieldTypeURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		FieldTypes []models.AppFWFieldType `json:"appfwfieldtype"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.FieldTypes, nil
}

func (s *AppFWService) GetAppFWFieldType(name string) (*models.AppFWFieldType, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWFieldTypeURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		FieldTypes []models.AppFWFieldType `json:"appfwfieldtype"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.FieldTypes) == 0 {
		return nil, fmt.Errorf("appfwfieldtype not found")
	}

	return &result.FieldTypes[0], nil
}

func (s *AppFWService) CountAppFWFieldType() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWFieldTypeURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		FieldTypes []struct {
			Count float64 `json:"__count"`
		} `json:"appfwfieldtype"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.FieldTypes) > 0 {
		return result.FieldTypes[0].Count, nil
	}

	return 0, nil
}

// appfwglobal_appfwpolicy_binding
// Binding object showing the appfwpolicy that can be bound to appfwglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwglobal_appfwpolicy_binding
func (s *AppFWService) AddAppFWGlobalAppFWPolicyBinding(resource models.AppFWGlobalAppFWPolicyBinding) error {
	payload := map[string]any{
		"appfwglobal_appfwpolicy_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWGlobalAppFWPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWGlobalAppFWPolicyBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWGlobalAppFWPolicyBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAppFWGlobalAppFWPolicyBinding() ([]models.AppFWGlobalAppFWPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWGlobalAppFWPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWGlobalAppFWPolicyBinding `json:"appfwglobal_appfwpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWGlobalAppFWPolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWGlobalAppFWPolicyBindingURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwglobal_appfwpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwglobal_auditnslogpolicy_binding
// Binding object showing the auditnslogpolicy that can be bound to appfwglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwglobal_auditnslogpolicy_binding
func (s *AppFWService) AddAppFWGlobalAuditNSLogPolicyBinding(resource models.AppFWGlobalAuditNSLogPolicyBinding) error {
	payload := map[string]any{
		"appfwglobal_auditnslogpolicy_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWGlobalAuditNSLogPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWGlobalAuditNSLogPolicyBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWGlobalAuditNSLogPolicyBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAppFWGlobalAuditNSLogPolicyBinding() ([]models.AppFWGlobalAuditNSLogPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWGlobalAuditNSLogPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWGlobalAuditNSLogPolicyBinding `json:"appfwglobal_auditnslogpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWGlobalAuditNSLogPolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWGlobalAuditNSLogPolicyBindingURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwglobal_auditnslogpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwglobal_audtisyslogpolicy_binding
// Binding object showing the auditsyslogpolicy that can be bound to appfwglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwglobal_auditsyslogpolicy_binding
func (s *AppFWService) AddAppFWGlobalAuditSyslogPolicyBinding(resource models.AppFWGlobalAuditSyslogPolicyBinding) error {
	payload := map[string]any{
		"appfwglobal_auditsyslogpolicy_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWGlobalAuditSyslogPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWGlobalAuditSyslogPolicyBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWGlobalAuditSyslogPolicyBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAppFWGlobalAuditSyslogPolicyBinding() ([]models.AppFWGlobalAuditSyslogPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWGlobalAuditSyslogPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWGlobalAuditSyslogPolicyBinding `json:"appfwglobal_auditsyslogpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWGlobalAuditSyslogPolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWGlobalAuditSyslogPolicyBindingURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwglobal_auditsyslogpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwglobal_binding
// Binding object which returns the resources bound to appfwglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwglobal_binding
func (s *AppFWService) GetAppFWGlobalBinding() (*models.AppFWGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWGlobalBinding `json:"appfwglobal_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) == 0 {
		return nil, fmt.Errorf("appfwglobal_binding not found")
	}

	return &result.Bindings[0], nil
}

// appfwhtmlerrorpage
// Configuration for HTML error page resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwhtmlerrorpage
func (s *AppFWService) DeleteAppFWHTMLErrorPage(name string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWHTMLErrorPageURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWHTMLErrorPage() ([]models.AppFWHTMLErrorPage, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWHTMLErrorPageURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Pages []models.AppFWHTMLErrorPage `json:"appfwhtmlerrorpage"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Pages, nil
}

func (s *AppFWService) GetAppFWHTMLErrorPage(name string) (*models.AppFWHTMLErrorPage, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWHTMLErrorPageURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Pages []models.AppFWHTMLErrorPage `json:"appfwhtmlerrorpage"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Pages) == 0 {
		return nil, fmt.Errorf("appfwhtmlerrorpage not found")
	}

	return &result.Pages[0], nil
}

func (s *AppFWService) ImportAppFWHTMLErrorPage(resource models.AppFWHTMLErrorPage) error {
	payload := map[string]any{
		"appfwhtmlerrorpage": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWHTMLErrorPageURL+"?action=Import", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) ChangeAppFWHTMLErrorPage(resource models.AppFWHTMLErrorPage) error {
	payload := map[string]any{
		"appfwhtmlerrorpage": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWHTMLErrorPageURL+"?action=change", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// appfwjsoncontenttype
// Configuration for JSON content type resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwjsoncontenttype
func (s *AppFWService) AddAppFWJSONContentType(resource models.AppFWJSONContentType) error {
	payload := map[string]any{
		"appfwjsoncontenttype": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWJSONContentTypeURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWJSONContentType(name string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWJSONContentTypeURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWJSONContentType() ([]models.AppFWJSONContentType, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWJSONContentTypeURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Types []models.AppFWJSONContentType `json:"appfwjsoncontenttype"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Types, nil
}

func (s *AppFWService) GetAppFWJSONContentType(name string) (*models.AppFWJSONContentType, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWJSONContentTypeURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Types []models.AppFWJSONContentType `json:"appfwjsoncontenttype"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Types) == 0 {
		return nil, fmt.Errorf("appfwjsoncontenttype not found")
	}

	return &result.Types[0], nil
}

func (s *AppFWService) CountAppFWJSONContentType() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWJSONContentTypeURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Types []struct {
			Count float64 `json:"__count"`
		} `json:"appfwjsoncontenttype"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Types) > 0 {
		return result.Types[0].Count, nil
	}

	return 0, nil
}

// appfwjsonerrorpage
// Configuration for JSON error page resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwjsonerrorpage
func (s *AppFWService) DeleteAppFWJSONErrorPage(name string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWJSONErrorPageURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWJSONErrorPage() ([]models.AppFWJSONErrorPage, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWJSONErrorPageURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Pages []models.AppFWJSONErrorPage `json:"appfwjsonerrorpage"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Pages, nil
}

func (s *AppFWService) GetAppFWJSONErrorPage(name string) (*models.AppFWJSONErrorPage, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWJSONErrorPageURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Pages []models.AppFWJSONErrorPage `json:"appfwjsonerrorpage"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Pages) == 0 {
		return nil, fmt.Errorf("appfwjsonerrorpage not found")
	}

	return &result.Pages[0], nil
}

func (s *AppFWService) ImportAppFWJSONErrorPage(resource models.AppFWJSONErrorPage) error {
	payload := map[string]any{
		"appfwjsonerrorpage": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWJSONErrorPageURL+"?action=Import", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) ChangeAppFWJSONErrorPage(resource models.AppFWJSONErrorPage) error {
	payload := map[string]any{
		"appfwjsonerrorpage": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWJSONErrorPageURL+"?action=change", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// appfwlearningdata
// Configuration for learning data resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwlearningdata
func (s *AppFWService) DeleteAppFWLearningData(name string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWLearningDataURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWLearningData() ([]models.AppFWLearningData, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWLearningDataURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AppFWLearningData `json:"appfwlearningdata"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AppFWService) CountAppFWLearningData() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWLearningDataURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []struct {
			Count float64 `json:"__count"`
		} `json:"appfwlearningdata"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) > 0 {
		return result.Data[0].Count, nil
	}

	return 0, nil
}

func (s *AppFWService) ResetAppFWLearningData(resource models.AppFWLearningData) error {
	payload := map[string]any{
		"appfwlearningdata": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWLearningDataURL+"?action=reset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) ExportAppFWLearningData(resource models.AppFWLearningData) error {
	payload := map[string]any{
		"appfwlearningdata": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWLearningDataURL+"?action=export", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// appfwlearningsettings
// Configuration for learning settings resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwlearningsettings
func (s *AppFWService) UpdateAppFWLearningSettings(resource models.AppFWLearningSettings) error {
	payload := map[string]any{
		"appfwlearningsettings": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, appFWLearningSettingsURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) UnsetAppFWLearningSettings(resource models.AppFWLearningSettings) error {
	payload := map[string]any{
		"appfwlearningsettings": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWLearningSettingsURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWLearningSettings() ([]models.AppFWLearningSettings, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWLearningSettingsURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Settings []models.AppFWLearningSettings `json:"appfwlearningsettings"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Settings, nil
}

func (s *AppFWService) GetAppFWLearningSettings(name string) (*models.AppFWLearningSettings, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWLearningSettingsURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Settings []models.AppFWLearningSettings `json:"appfwlearningsettings"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Settings) == 0 {
		return nil, fmt.Errorf("appfwlearningsettings not found")
	}

	return &result.Settings[0], nil
}

func (s *AppFWService) CountAppFWLearningSettings() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWLearningSettingsURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Settings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwlearningsettings"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Settings) > 0 {
		return result.Settings[0].Count, nil
	}

	return 0, nil
}

// appfwmultipartformcontenttype
// Configuration for Multipart form content type resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwmultipartformcontenttype
func (s *AppFWService) AddAppFWMultipartFormContentType(resource models.AppFWMultipartFormContentType) error {
	payload := map[string]any{
		"appfwmultipartformcontenttype": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWMultipartFormContentTypeURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWMultipartFormContentType(name string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWMultipartFormContentTypeURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWMultipartFormContentType() ([]models.AppFWMultipartFormContentType, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWMultipartFormContentTypeURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Types []models.AppFWMultipartFormContentType `json:"appfwmultipartformcontenttype"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Types, nil
}

func (s *AppFWService) GetAppFWMultipartFormContentType(name string) (*models.AppFWMultipartFormContentType, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWMultipartFormContentTypeURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Types []models.AppFWMultipartFormContentType `json:"appfwmultipartformcontenttype"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Types) == 0 {
		return nil, fmt.Errorf("appfwmultipartformcontenttype not found")
	}

	return &result.Types[0], nil
}

func (s *AppFWService) CountAppFWMultipartFormContentType() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWMultipartFormContentTypeURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Types []struct {
			Count float64 `json:"__count"`
		} `json:"appfwmultipartformcontenttype"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Types) > 0 {
		return result.Types[0].Count, nil
	}

	return 0, nil
}

// appfwpolicy
// Configuration for application firewall policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwpolicy
func (s *AppFWService) AddAppFWPolicy(resource models.AppFWPolicy) error {
	payload := map[string]any{
		"appfwpolicy": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWPolicy(name string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWPolicyURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) UpdateAppFWPolicy(resource models.AppFWPolicy) error {
	payload := map[string]any{
		"appfwpolicy": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, appFWPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) UnsetAppFWPolicy(resource models.AppFWPolicy) error {
	payload := map[string]any{
		"appfwpolicy": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWPolicyURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWPolicy() ([]models.AppFWPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWPolicyURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Policies []models.AppFWPolicy `json:"appfwpolicy"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Policies, nil
}

func (s *AppFWService) GetAppFWPolicy(name string) (*models.AppFWPolicy, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWPolicyURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Policies []models.AppFWPolicy `json:"appfwpolicy"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Policies) == 0 {
		return nil, fmt.Errorf("appfwpolicy not found")
	}

	return &result.Policies[0], nil
}

func (s *AppFWService) CountAppFWPolicy() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWPolicyURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Policies []struct {
			Count float64 `json:"__count"`
		} `json:"appfwpolicy"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Policies) > 0 {
		return result.Policies[0].Count, nil
	}

	return 0, nil
}

func (s *AppFWService) RenameAppFWPolicy(resource models.AppFWPolicy) error {
	payload := map[string]any{
		"appfwpolicy": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWPolicyURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// appfwpolicylabel
// Configuration for application firewall policy label resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwpolicylabel
func (s *AppFWService) AddAppFWPolicyLabel(resource models.AppFWPolicyLabel) error {
	payload := map[string]any{
		"appfwpolicylabel": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWPolicyLabelURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWPolicyLabel(name string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWPolicyLabelURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWPolicyLabel() ([]models.AppFWPolicyLabel, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWPolicyLabelURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Labels []models.AppFWPolicyLabel `json:"appfwpolicylabel"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Labels, nil
}

func (s *AppFWService) GetAppFWPolicyLabel(name string) (*models.AppFWPolicyLabel, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWPolicyLabelURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Labels []models.AppFWPolicyLabel `json:"appfwpolicylabel"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Labels) == 0 {
		return nil, fmt.Errorf("appfwpolicylabel not found")
	}

	return &result.Labels[0], nil
}

func (s *AppFWService) CountAppFWPolicyLabel() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWPolicyLabelURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Labels []struct {
			Count float64 `json:"__count"`
		} `json:"appfwpolicylabel"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Labels) > 0 {
		return result.Labels[0].Count, nil
	}

	return 0, nil
}

func (s *AppFWService) RenameAppFWPolicyLabel(resource models.AppFWPolicyLabel) error {
	payload := map[string]any{
		"appfwpolicylabel": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWPolicyLabelURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// appfwpolicylabel_appfwpolicy_binding
// Binding object showing the appfwpolicy that can be bound to appfwpolicylabel.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwpolicylabel_appfwpolicy_binding
func (s *AppFWService) AddAppFWPolicyLabelAppFWPolicyBinding(resource models.AppFWPolicyLabelAppFWPolicyBinding) error {
	payload := map[string]any{
		"appfwpolicylabel_appfwpolicy_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWPolicyLabelAppFWPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWPolicyLabelAppFWPolicyBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWPolicyLabelAppFWPolicyBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWPolicyLabelAppFWPolicyBinding() ([]models.AppFWPolicyLabelAppFWPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWPolicyLabelAppFWPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWPolicyLabelAppFWPolicyBinding `json:"appfwpolicylabel_appfwpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWPolicyLabelAppFWPolicyBinding(name string) ([]models.AppFWPolicyLabelAppFWPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWPolicyLabelAppFWPolicyBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWPolicyLabelAppFWPolicyBinding `json:"appfwpolicylabel_appfwpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWPolicyLabelAppFWPolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFWPolicyLabelAppFWPolicyBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwpolicylabel_appfwpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwpolicylabel_binding
// Binding object which returns the resources bound to appfwpolicylabel.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwpolicylabel_binding
func (s *AppFWService) GetAllAppFWPolicyLabelBinding() ([]models.AppFWPolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWPolicyLabelBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWPolicyLabelBinding `json:"appfwpolicylabel_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWPolicyLabelBinding(name string) (*models.AppFWPolicyLabelBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWPolicyLabelBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWPolicyLabelBinding `json:"appfwpolicylabel_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) == 0 {
		return nil, fmt.Errorf("appfwpolicylabel_binding not found")
	}

	return &result.Bindings[0], nil
}

// appfwpolicylabel_policybinding_binding
// Binding object showing the policybinding that can be bound to appfwpolicylabel.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwpolicylabel_policybinding_binding
func (s *AppFWService) GetAllAppFWPolicyLabelPolicyBindingBinding() ([]models.AppFWPolicyLabelPolicyBindingBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWPolicyLabelPolicyBindingBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWPolicyLabelPolicyBindingBinding `json:"appfwpolicylabel_policybinding_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWPolicyLabelPolicyBindingBinding(name string) ([]models.AppFWPolicyLabelPolicyBindingBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWPolicyLabelPolicyBindingBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWPolicyLabelPolicyBindingBinding `json:"appfwpolicylabel_policybinding_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWPolicyLabelPolicyBindingBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFWPolicyLabelPolicyBindingBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwpolicylabel_policybinding_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwpolicy_appfwglobal_binding
// Binding object showing the appfwglobal that can be bound to appfwpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwpolicy_appfwglobal_binding
func (s *AppFWService) GetAllAppFWPolicyAppFWGlobalBinding() ([]models.AppFWPolicyAppFWGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWPolicyAppFWGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWPolicyAppFWGlobalBinding `json:"appfwpolicy_appfwglobal_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWPolicyAppFWGlobalBinding(name string) ([]models.AppFWPolicyAppFWGlobalBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWPolicyAppFWGlobalBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWPolicyAppFWGlobalBinding `json:"appfwpolicy_appfwglobal_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWPolicyAppFWGlobalBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFWPolicyAppFWGlobalBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwpolicy_appfwglobal_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwpolicy_appfwpolicylabel_binding
// Binding object showing the appfwpolicylabel that can be bound to appfwpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwpolicy_appfwpolicylabel_binding
func (s *AppFWService) GetAllAppFWPolicyAppFWPolicyLabelBinding() ([]models.AppFWPolicyAppFWPolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWPolicyAppFWPolicyLabelBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWPolicyAppFWPolicyLabelBinding `json:"appfwpolicy_appfwpolicylabel_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWPolicyAppFWPolicyLabelBinding(name string) ([]models.AppFWPolicyAppFWPolicyLabelBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWPolicyAppFWPolicyLabelBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWPolicyAppFWPolicyLabelBinding `json:"appfwpolicy_appfwpolicylabel_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWPolicyAppFWPolicyLabelBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFWPolicyAppFWPolicyLabelBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwpolicy_appfwpolicylabel_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwpolicy_binding
// Binding object which returns the resources bound to appfwpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwpolicy_binding
func (s *AppFWService) GetAllAppFWPolicyBinding() ([]models.AppFWPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWPolicyBinding `json:"appfwpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWPolicyBinding(name string) (*models.AppFWPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWPolicyBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWPolicyBinding `json:"appfwpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) == 0 {
		return nil, fmt.Errorf("appfwpolicy_binding not found")
	}

	return &result.Bindings[0], nil
}

// appfwpolicy_csvserver_binding
// Binding object showing the csvserver that can be bound to appfwpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwpolicy_csvserver_binding
func (s *AppFWService) GetAllAppFWPolicyCSVServerBinding() ([]models.AppFWPolicyCSVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWPolicyCSVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWPolicyCSVServerBinding `json:"appfwpolicy_csvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWPolicyCSVServerBinding(name string) ([]models.AppFWPolicyCSVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWPolicyCSVServerBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWPolicyCSVServerBinding `json:"appfwpolicy_csvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWPolicyCSVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFWPolicyCSVServerBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwpolicy_csvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwpolicy_lbvserver_binding
// Binding object showing the lbvserver that can be bound to appfwpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwpolicy_lbvserver_binding
func (s *AppFWService) GetAllAppFWPolicyLBVServerBinding() ([]models.AppFWPolicyLBVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWPolicyLBVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWPolicyLBVServerBinding `json:"appfwpolicy_lbvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWPolicyLBVServerBinding(name string) ([]models.AppFWPolicyLBVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWPolicyLBVServerBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWPolicyLBVServerBinding `json:"appfwpolicy_lbvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWPolicyLBVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFWPolicyLBVServerBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwpolicy_lbvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwprofile
// Configuration for application firewall profile resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile
func (s *AppFWService) AddAppFWProfile(resource models.AppFWProfile) error {
	payload := map[string]interface{}{
		"appfwprofile": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWProfile(name string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) UpdateAppFWProfile(resource models.AppFWProfile) error {
	payload := map[string]interface{}{
		"appfwprofile": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, appFWProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) UnsetAppFWProfile(resource models.AppFWProfile) error {
	payload := map[string]interface{}{
		"appfwprofile": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWProfileURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWProfile() ([]models.AppFWProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Profiles []models.AppFWProfile `json:"appfwprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Profiles, nil
}

func (s *AppFWService) GetAppFWProfile(name string) (*models.AppFWProfile, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Profiles []models.AppFWProfile `json:"appfwprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Profiles) == 0 {
		return nil, fmt.Errorf("appfwprofile not found")
	}

	return &result.Profiles[0], nil
}

func (s *AppFWService) CountAppFWProfile() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWProfileURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Profiles []struct {
			Count float64 `json:"__count"`
		} `json:"appfwprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Profiles) > 0 {
		return result.Profiles[0].Count, nil
	}

	return 0, nil
}

func (s *AppFWService) RestoreAppFWProfile(resource models.AppFWProfile) error {
	payload := map[string]interface{}{
		"appfwprofile": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWProfileURL+"?action=restore", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// appfwprofile_binding
// Binding object which returns the resources bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_binding
func (s *AppFWService) GetAllAppFWProfileBinding() ([]models.AppFWProfileBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWProfileBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileBinding `json:"appfwprofile_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWProfileBinding(name string) (*models.AppFWProfileBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileBinding `json:"appfwprofile_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) == 0 {
		return nil, fmt.Errorf("appfwprofile_binding not found")
	}

	return &result.Bindings[0], nil
}

// appfwprofile_cmdinjection_binding
// Binding object showing the cmdinjection that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_cmdinjection_binding
func (s *AppFWService) AddAppFWProfileCMDInjectionBinding(resource models.AppFWProfileCMDInjectionBinding) error {
	payload := map[string]interface{}{
		"appfwprofile_cmdinjection_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWProfileCMDInjectionBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWProfileCMDInjectionBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileCMDInjectionBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWProfileCMDInjectionBinding() ([]models.AppFWProfileCMDInjectionBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWProfileCMDInjectionBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileCMDInjectionBinding `json:"appfwprofile_cmdinjection_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWProfileCMDInjectionBinding(name string) ([]models.AppFWProfileCMDInjectionBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileCMDInjectionBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileCMDInjectionBinding `json:"appfwprofile_cmdinjection_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWProfileCMDInjectionBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFWProfileCMDInjectionBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwprofile_cmdinjection_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwprofile_contenttype_binding
// Binding object showing the contenttype that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_contenttype_binding
func (s *AppFWService) AddAppFWProfileContentTypeBinding(resource models.AppFWProfileContentTypeBinding) error {
	payload := map[string]interface{}{
		"appfwprofile_contenttype_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWProfileContentTypeBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWProfileContentTypeBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileContentTypeBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWProfileContentTypeBinding() ([]models.AppFWProfileContentTypeBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWProfileContentTypeBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileContentTypeBinding `json:"appfwprofile_contenttype_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWProfileContentTypeBinding(name string) ([]models.AppFWProfileContentTypeBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileContentTypeBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileContentTypeBinding `json:"appfwprofile_contenttype_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWProfileContentTypeBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFWProfileContentTypeBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwprofile_contenttype_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwprofile_cookieconsistency_binding
// Binding object showing the cookieconsistency that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_cookieconsistency_binding
func (s *AppFWService) AddAppFWProfileCookieConsistencyBinding(resource models.AppFWProfileCookieConsistencyBinding) error {
	payload := map[string]interface{}{
		"appfwprofile_cookieconsistency_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWProfileCookieConsistencyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWProfileCookieConsistencyBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileCookieConsistencyBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWProfileCookieConsistencyBinding() ([]models.AppFWProfileCookieConsistencyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWProfileCookieConsistencyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileCookieConsistencyBinding `json:"appfwprofile_cookieconsistency_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWProfileCookieConsistencyBinding(name string) ([]models.AppFWProfileCookieConsistencyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileCookieConsistencyBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileCookieConsistencyBinding `json:"appfwprofile_cookieconsistency_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWProfileCookieConsistencyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFWProfileCookieConsistencyBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwprofile_cookieconsistency_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwprofile_creditcardnumber_binding
// Binding object showing the creditcardnumber that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_creditcardnumber_binding
func (s *AppFWService) AddAppFWProfileCreditCardNumberBinding(resource models.AppFWProfileCreditCardNumberBinding) error {
	payload := map[string]interface{}{
		"appfwprofile_creditcardnumber_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWProfileCreditCardNumberBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWProfileCreditCardNumberBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileCreditCardNumberBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWProfileCreditCardNumberBinding() ([]models.AppFWProfileCreditCardNumberBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWProfileCreditCardNumberBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileCreditCardNumberBinding `json:"appfwprofile_creditcardnumber_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWProfileCreditCardNumberBinding(name string) ([]models.AppFWProfileCreditCardNumberBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileCreditCardNumberBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileCreditCardNumberBinding `json:"appfwprofile_creditcardnumber_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWProfileCreditCardNumberBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFWProfileCreditCardNumberBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwprofile_creditcardnumber_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwprofile_crosssitescripting_binding
// Binding object showing the crosssitescripting that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_crosssitescripting_binding
func (s *AppFWService) AddAppFWProfileCrossSiteScriptingBinding(resource models.AppFWProfileCrossSiteScriptingBinding) error {
	payload := map[string]interface{}{
		"appfwprofile_crosssitescripting_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWProfileCrossSiteScriptingBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWProfileCrossSiteScriptingBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileCrossSiteScriptingBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWProfileCrossSiteScriptingBinding() ([]models.AppFWProfileCrossSiteScriptingBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWProfileCrossSiteScriptingBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileCrossSiteScriptingBinding `json:"appfwprofile_crosssitescripting_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWProfileCrossSiteScriptingBinding(name string) ([]models.AppFWProfileCrossSiteScriptingBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileCrossSiteScriptingBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileCrossSiteScriptingBinding `json:"appfwprofile_crosssitescripting_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWProfileCrossSiteScriptingBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFWProfileCrossSiteScriptingBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwprofile_crosssitescripting_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwprofile_csrftag_binding
// Binding object showing the csrftag that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_csrftag_binding
func (s *AppFWService) AddAppFWProfileCSRFTagBinding(resource models.AppFWProfileCSRFTagBinding) error {
	payload := map[string]interface{}{
		"appfwprofile_csrftag_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWProfileCSRFTagBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWProfileCSRFTagBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileCSRFTagBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWProfileCSRFTagBinding() ([]models.AppFWProfileCSRFTagBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWProfileCSRFTagBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileCSRFTagBinding `json:"appfwprofile_csrftag_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWProfileCSRFTagBinding(name string) ([]models.AppFWProfileCSRFTagBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileCSRFTagBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileCSRFTagBinding `json:"appfwprofile_csrftag_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWProfileCSRFTagBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFWProfileCSRFTagBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwprofile_csrftag_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwprofile_denyurl_binding
// Binding object showing the denyurl that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_denyurl_binding
func (s *AppFWService) AddAppFWProfileDenyURLBinding(resource models.AppFWProfileDenyURLBinding) error {
	payload := map[string]interface{}{
		"appfwprofile_denyurl_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWProfileDenyURLBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWProfileDenyURLBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileDenyURLBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWProfileDenyURLBinding() ([]models.AppFWProfileDenyURLBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWProfileDenyURLBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileDenyURLBinding `json:"appfwprofile_denyurl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWProfileDenyURLBinding(name string) ([]models.AppFWProfileDenyURLBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileDenyURLBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileDenyURLBinding `json:"appfwprofile_denyurl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWProfileDenyURLBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFWProfileDenyURLBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwprofile_denyurl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwprofile_excluderescontenttype_binding
// Binding object showing the excluderescontenttype that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_excluderescontenttype_binding
func (s *AppFWService) AddAppFWProfileExcludeRESContentTypeBinding(resource models.AppFWProfileExcludeRESContentTypeBinding) error {
	payload := map[string]interface{}{
		"appfwprofile_excluderescontenttype_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWProfileExcludeRESContentTypeBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWProfileExcludeRESContentTypeBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileExcludeRESContentTypeBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWProfileExcludeRESContentTypeBinding() ([]models.AppFWProfileExcludeRESContentTypeBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWProfileExcludeRESContentTypeBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileExcludeRESContentTypeBinding `json:"appfwprofile_excluderescontenttype_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWProfileExcludeRESContentTypeBinding(name string) ([]models.AppFWProfileExcludeRESContentTypeBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileExcludeRESContentTypeBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileExcludeRESContentTypeBinding `json:"appfwprofile_excluderescontenttype_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWProfileExcludeRESContentTypeBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFWProfileExcludeRESContentTypeBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwprofile_excluderescontenttype_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwprofile_fieldconsistency_binding
// Binding object showing the fieldconsistency that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_fieldconsistency_binding
func (s *AppFWService) AddAppFWProfileFieldConsistencyBinding(resource models.AppFWProfileFieldConsistencyBinding) error {
	payload := map[string]interface{}{
		"appfwprofile_fieldconsistency_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWProfileFieldConsistencyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWProfileFieldConsistencyBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileFieldConsistencyBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWProfileFieldConsistencyBinding() ([]models.AppFWProfileFieldConsistencyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWProfileFieldConsistencyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileFieldConsistencyBinding `json:"appfwprofile_fieldconsistency_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWProfileFieldConsistencyBinding(name string) ([]models.AppFWProfileFieldConsistencyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileFieldConsistencyBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileFieldConsistencyBinding `json:"appfwprofile_fieldconsistency_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWProfileFieldConsistencyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFWProfileFieldConsistencyBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwprofile_fieldconsistency_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwprofile_fieldformat_binding
// Binding object showing the fieldformat that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_fieldformat_binding
func (s *AppFWService) AddAppFWProfileFieldFormatBinding(resource models.AppFWProfileFieldFormatBinding) error {
	payload := map[string]interface{}{
		"appfwprofile_fieldformat_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWProfileFieldFormatBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWProfileFieldFormatBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileFieldFormatBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWProfileFieldFormatBinding() ([]models.AppFWProfileFieldFormatBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWProfileFieldFormatBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileFieldFormatBinding `json:"appfwprofile_fieldformat_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWProfileFieldFormatBinding(name string) ([]models.AppFWProfileFieldFormatBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileFieldFormatBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileFieldFormatBinding `json:"appfwprofile_fieldformat_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWProfileFieldFormatBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFWProfileFieldFormatBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwprofile_fieldformat_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwprofile_fileuploadtype_binding
// Binding object showing the fileuploadtype that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_fileuploadtype_binding
func (s *AppFWService) AddAppFWProfileFileUploadTypeBinding(resource models.AppFWProfileFileUploadTypeBinding) error {
	payload := map[string]interface{}{
		"appfwprofile_fileuploadtype_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWProfileFileUploadTypeBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWProfileFileUploadTypeBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileFileUploadTypeBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWProfileFileUploadTypeBinding() ([]models.AppFWProfileFileUploadTypeBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWProfileFileUploadTypeBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileFileUploadTypeBinding `json:"appfwprofile_fileuploadtype_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWProfileFileUploadTypeBinding(name string) ([]models.AppFWProfileFileUploadTypeBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileFileUploadTypeBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileFileUploadTypeBinding `json:"appfwprofile_fileuploadtype_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWProfileFileUploadTypeBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFWProfileFileUploadTypeBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwprofile_fileuploadtype_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwprofile_jsondosurl_binding
// Binding object showing the jsondosurl that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_jsondosurl_binding
func (s *AppFWService) AddAppFWProfileJSONDOSURLBinding(resource models.AppFWProfileJSONDOSURLBinding) error {
	payload := map[string]interface{}{
		"appfwprofile_jsondosurl_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWProfileJSONDOSURLBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWProfileJSONDOSURLBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileJSONDOSURLBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWProfileJSONDOSURLBinding() ([]models.AppFWProfileJSONDOSURLBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWProfileJSONDOSURLBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileJSONDOSURLBinding `json:"appfwprofile_jsondosurl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWProfileJSONDOSURLBinding(name string) ([]models.AppFWProfileJSONDOSURLBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileJSONDOSURLBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileJSONDOSURLBinding `json:"appfwprofile_jsondosurl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWProfileJSONDOSURLBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFWProfileJSONDOSURLBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwprofile_jsondosurl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwprofile_jsonsqlurl_binding
// Binding object showing the jsonsqlurl that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_jsonsqlurl_binding
func (s *AppFWService) AddAppFWProfileJSONSQLURLBinding(resource models.AppFWProfileJSONSQLURLBinding) error {
	payload := map[string]interface{}{
		"appfwprofile_jsonsqlurl_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWProfileJSONSQLURLBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWProfileJSONSQLURLBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileJSONSQLURLBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWProfileJSONSQLURLBinding() ([]models.AppFWProfileJSONSQLURLBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWProfileJSONSQLURLBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileJSONSQLURLBinding `json:"appfwprofile_jsonsqlurl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWProfileJSONSQLURLBinding(name string) ([]models.AppFWProfileJSONSQLURLBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileJSONSQLURLBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileJSONSQLURLBinding `json:"appfwprofile_jsonsqlurl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWProfileJSONSQLURLBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFWProfileJSONSQLURLBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwprofile_jsonsqlurl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwprofile_jsonxssurl_binding
// Binding object showing the jsonxssurl that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_jsonxssurl_binding
func (s *AppFWService) AddAppFWProfileJSONXSSURLBinding(resource models.AppFWProfileJSONXSSURLBinding) error {
	payload := map[string]interface{}{
		"appfwprofile_jsonxssurl_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWProfileJSONXSSURLBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWProfileJSONXSSURLBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileJSONXSSURLBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWProfileJSONXSSURLBinding() ([]models.AppFWProfileJSONXSSURLBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWProfileJSONXSSURLBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileJSONXSSURLBinding `json:"appfwprofile_jsonxssurl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWProfileJSONXSSURLBinding(name string) ([]models.AppFWProfileJSONXSSURLBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileJSONXSSURLBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileJSONXSSURLBinding `json:"appfwprofile_jsonxssurl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWProfileJSONXSSURLBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFWProfileJSONXSSURLBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwprofile_jsonxssurl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwprofile_logexpression_binding
// Binding object showing the logexpression that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_logexpression_binding
func (s *AppFWService) AddAppFWProfileLogExpressionBinding(resource models.AppFWProfileLogExpressionBinding) error {
	payload := map[string]interface{}{
		"appfwprofile_logexpression_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWProfileLogExpressionBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWProfileLogExpressionBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileLogExpressionBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWProfileLogExpressionBinding() ([]models.AppFWProfileLogExpressionBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWProfileLogExpressionBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileLogExpressionBinding `json:"appfwprofile_logexpression_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWProfileLogExpressionBinding(name string) ([]models.AppFWProfileLogExpressionBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileLogExpressionBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileLogExpressionBinding `json:"appfwprofile_logexpression_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWProfileLogExpressionBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFWProfileLogExpressionBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwprofile_logexpression_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwprofile_safeobject_binding
// Binding object showing the safeobject that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_safeobject_binding
func (s *AppFWService) AddAppFWProfileSafeObjectBinding(resource models.AppFWProfileSafeObjectBinding) error {
	payload := map[string]interface{}{
		"appfwprofile_safeobject_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWProfileSafeObjectBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWProfileSafeObjectBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileSafeObjectBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWProfileSafeObjectBinding() ([]models.AppFWProfileSafeObjectBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWProfileSafeObjectBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileSafeObjectBinding `json:"appfwprofile_safeobject_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWProfileSafeObjectBinding(name string) ([]models.AppFWProfileSafeObjectBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileSafeObjectBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileSafeObjectBinding `json:"appfwprofile_safeobject_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWProfileSafeObjectBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFWProfileSafeObjectBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwprofile_safeobject_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwprofile_sqlinjection_binding
// Binding object showing the sqlinjection that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_sqlinjection_binding
func (s *AppFWService) AddAppFWProfileSQLInjectionBinding(resource models.AppFWProfileSQLInjectionBinding) error {
	payload := map[string]interface{}{
		"appfwprofile_sqlinjection_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWProfileSQLInjectionBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWProfileSQLInjectionBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileSQLInjectionBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWProfileSQLInjectionBinding() ([]models.AppFWProfileSQLInjectionBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWProfileSQLInjectionBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileSQLInjectionBinding `json:"appfwprofile_sqlinjection_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWProfileSQLInjectionBinding(name string) ([]models.AppFWProfileSQLInjectionBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileSQLInjectionBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileSQLInjectionBinding `json:"appfwprofile_sqlinjection_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWProfileSQLInjectionBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFWProfileSQLInjectionBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwprofile_sqlinjection_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwprofile_starturl_binding
// Binding object showing the starturl that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_starturl_binding
func (s *AppFWService) AddAppFWProfileStartURLBinding(resource models.AppFWProfileStartURLBinding) error {
	payload := map[string]interface{}{
		"appfwprofile_starturl_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWProfileStartURLBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWProfileStartURLBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileStartURLBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWProfileStartURLBinding() ([]models.AppFWProfileStartURLBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWProfileStartURLBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileStartURLBinding `json:"appfwprofile_starturl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWProfileStartURLBinding(name string) ([]models.AppFWProfileStartURLBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileStartURLBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileStartURLBinding `json:"appfwprofile_starturl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWProfileStartURLBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFWProfileStartURLBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwprofile_starturl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwprofile_trustedlearningclients_binding
// Binding object showing the trustedlearningclients that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_trustedlearningclients_binding
func (s *AppFWService) AddAppFWProfileTrustedLearningClientsBinding(resource models.AppFWProfileTrustedLearningClientsBinding) error {
	payload := map[string]interface{}{
		"appfwprofile_trustedlearningclients_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWProfileTrustedLearningClientsBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWProfileTrustedLearningClientsBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileTrustedLearningClientsBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWProfileTrustedLearningClientsBinding() ([]models.AppFWProfileTrustedLearningClientsBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWProfileTrustedLearningClientsBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileTrustedLearningClientsBinding `json:"appfwprofile_trustedlearningclients_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWProfileTrustedLearningClientsBinding(name string) ([]models.AppFWProfileTrustedLearningClientsBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileTrustedLearningClientsBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileTrustedLearningClientsBinding `json:"appfwprofile_trustedlearningclients_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWProfileTrustedLearningClientsBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFWProfileTrustedLearningClientsBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwprofile_trustedlearningclients_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwprofile_xmlattachmenturl_binding
// Binding object showing the xmlattachmenturl that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_xmlattachmenturl_binding
func (s *AppFWService) AddAppFWProfileXMLAttachmentURLBinding(resource models.AppFWProfileXMLAttachmentURLBinding) error {
	payload := map[string]interface{}{
		"appfwprofile_xmlattachmenturl_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWProfileXMLAttachmentURLBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWProfileXMLAttachmentURLBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileXMLAttachmentURLBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWProfileXMLAttachmentURLBinding() ([]models.AppFWProfileXMLAttachmentURLBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWProfileXMLAttachmentURLBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileXMLAttachmentURLBinding `json:"appfwprofile_xmlattachmenturl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWProfileXMLAttachmentURLBinding(name string) ([]models.AppFWProfileXMLAttachmentURLBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileXMLAttachmentURLBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileXMLAttachmentURLBinding `json:"appfwprofile_xmlattachmenturl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWProfileXMLAttachmentURLBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFWProfileXMLAttachmentURLBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwprofile_xmlattachmenturl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwprofile_xmldosurl_binding
// Binding object showing the xmldosurl that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_xmldosurl_binding
func (s *AppFWService) AddAppFWProfileXMLDOSURLBinding(resource models.AppFWProfileXMLDOSURLBinding) error {
	payload := map[string]interface{}{
		"appfwprofile_xmldosurl_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWProfileXMLDOSURLBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWProfileXMLDOSURLBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileXMLDOSURLBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWProfileXMLDOSURLBinding() ([]models.AppFWProfileXMLDOSURLBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWProfileXMLDOSURLBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileXMLDOSURLBinding `json:"appfwprofile_xmldosurl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWProfileXMLDOSURLBinding(name string) ([]models.AppFWProfileXMLDOSURLBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileXMLDOSURLBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileXMLDOSURLBinding `json:"appfwprofile_xmldosurl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWProfileXMLDOSURLBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFWProfileXMLDOSURLBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwprofile_xmldosurl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwprofile_xmlsqlinjection_binding
// Binding object showing the xmlsqlinjection that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_xmlsqlinjection_binding
func (s *AppFWService) AddAppFWProfileXMLSQLInjectionBinding(resource models.AppFWProfileXMLSQLInjectionBinding) error {
	payload := map[string]interface{}{
		"appfwprofile_xmlsqlinjection_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWProfileXMLSQLInjectionBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWProfileXMLSQLInjectionBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileXMLSQLInjectionBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWProfileXMLSQLInjectionBinding() ([]models.AppFWProfileXMLSQLInjectionBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWProfileXMLSQLInjectionBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileXMLSQLInjectionBinding `json:"appfwprofile_xmlsqlinjection_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWProfileXMLSQLInjectionBinding(name string) ([]models.AppFWProfileXMLSQLInjectionBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileXMLSQLInjectionBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileXMLSQLInjectionBinding `json:"appfwprofile_xmlsqlinjection_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWProfileXMLSQLInjectionBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFWProfileXMLSQLInjectionBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwprofile_xmlsqlinjection_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwprofile_xmlvalidationurl_binding
// Binding object showing the xmlvalidationurl that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_xmlvalidationurl_binding
func (s *AppFWService) AddAppFWProfileXMLValidationURLBinding(resource models.AppFWProfileXMLValidationURLBinding) error {
	payload := map[string]interface{}{
		"appfwprofile_xmlvalidationurl_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWProfileXMLValidationURLBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWProfileXMLValidationURLBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileXMLValidationURLBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWProfileXMLValidationURLBinding() ([]models.AppFWProfileXMLValidationURLBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWProfileXMLValidationURLBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileXMLValidationURLBinding `json:"appfwprofile_xmlvalidationurl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWProfileXMLValidationURLBinding(name string) ([]models.AppFWProfileXMLValidationURLBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileXMLValidationURLBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileXMLValidationURLBinding `json:"appfwprofile_xmlvalidationurl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWProfileXMLValidationURLBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFWProfileXMLValidationURLBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwprofile_xmlvalidationurl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwprofile_xmlwsiurl_binding
// Binding object showing the xmlwsiurl that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_xmlwsiurl_binding
func (s *AppFWService) AddAppFWProfileXMLWSIURLBinding(resource models.AppFWProfileXMLWSIURLBinding) error {
	payload := map[string]interface{}{
		"appfwprofile_xmlwsiurl_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWProfileXMLWSIURLBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWProfileXMLWSIURLBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileXMLWSIURLBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWProfileXMLWSIURLBinding() ([]models.AppFWProfileXMLWSIURLBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWProfileXMLWSIURLBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileXMLWSIURLBinding `json:"appfwprofile_xmlwsiurl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWProfileXMLWSIURLBinding(name string) ([]models.AppFWProfileXMLWSIURLBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileXMLWSIURLBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileXMLWSIURLBinding `json:"appfwprofile_xmlwsiurl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWProfileXMLWSIURLBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFWProfileXMLWSIURLBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwprofile_xmlwsiurl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwprofile_xmlxss_binding
// Binding object showing the xmlxss that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_xmlxss_binding
func (s *AppFWService) AddAppFWProfileXMLXSSBinding(resource models.AppFWProfileXMLXSSBinding) error {
	payload := map[string]interface{}{
		"appfwprofile_xmlxss_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWProfileXMLXSSBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWProfileXMLXSSBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileXMLXSSBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWProfileXMLXSSBinding() ([]models.AppFWProfileXMLXSSBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWProfileXMLXSSBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileXMLXSSBinding `json:"appfwprofile_xmlxss_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) GetAppFWProfileXMLXSSBinding(name string) ([]models.AppFWProfileXMLXSSBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWProfileXMLXSSBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppFWProfileXMLXSSBinding `json:"appfwprofile_xmlxss_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppFWService) CountAppFWProfileXMLXSSBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFWProfileXMLXSSBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"appfwprofile_xmlxss_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// appfwsettings
// Configuration for AS settings resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwsettings
func (s *AppFWService) UpdateAppFWSettings(resource models.AppFWSettings) error {
	payload := map[string]interface{}{
		"appfwsettings": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, appFWSettingsURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) UnsetAppFWSettings(resource models.AppFWSettings) error {
	payload := map[string]interface{}{
		"appfwsettings": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWSettingsURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWSettings() (*models.AppFWSettings, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWSettingsURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Settings []models.AppFWSettings `json:"appfwsettings"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Settings) == 0 {
		return nil, fmt.Errorf("appfwsettings not found")
	}

	return &result.Settings[0], nil
}

// appfwsignatures
// Configuration for application firewall signatures XML configuration resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwsignatures
func (s *AppFWService) DeleteAppFWSignatures(name string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWSignaturesURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWSignatures() ([]models.AppFWSignatures, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWSignaturesURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Signatures []models.AppFWSignatures `json:"appfwsignatures"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Signatures, nil
}

func (s *AppFWService) GetAppFWSignatures(name string) (*models.AppFWSignatures, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWSignaturesURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Signatures []models.AppFWSignatures `json:"appfwsignatures"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Signatures) == 0 {
		return nil, fmt.Errorf("appfwsignatures not found")
	}

	return &result.Signatures[0], nil
}

func (s *AppFWService) ImportAppFWSignatures(resource models.AppFWSignatures) error {
	payload := map[string]interface{}{
		"appfwsignatures": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWSignaturesURL+"?action=Import", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) ChangeAppFWSignatures(resource models.AppFWSignatures) error {
	payload := map[string]interface{}{
		"appfwsignatures": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWSignaturesURL+"?action=change", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// appfwtransactionrecords
// Configuration for Application firewall transaction record resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwtransactionrecords
func (s *AppFWService) GetAllAppFWTransactionRecords() ([]models.AppFWTransactionRecords, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWTransactionRecordsURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Records []models.AppFWTransactionRecords `json:"appfwtransactionrecords"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Records, nil
}

func (s *AppFWService) CountAppFWTransactionRecords() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWTransactionRecordsURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Records []struct {
			Count float64 `json:"__count"`
		} `json:"appfwtransactionrecords"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Records) > 0 {
		return result.Records[0].Count, nil
	}

	return 0, nil
}

// appfwurlencodedformcontenttype
// Configuration for Urlencoded form content type resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwurlencodedformcontenttype
func (s *AppFWService) AddAppFWURLEncodedFormContentType(resource models.AppFWURLEncodedFormContentType) error {
	payload := map[string]interface{}{
		"appfwurlencodedformcontenttype": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWURLEncodedFormContentTypeURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWURLEncodedFormContentType(name string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWURLEncodedFormContentTypeURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWURLEncodedFormContentType() ([]models.AppFWURLEncodedFormContentType, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWURLEncodedFormContentTypeURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Types []models.AppFWURLEncodedFormContentType `json:"appfwurlencodedformcontenttype"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Types, nil
}

func (s *AppFWService) GetAppFWURLEncodedFormContentType(name string) (*models.AppFWURLEncodedFormContentType, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWURLEncodedFormContentTypeURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Types []models.AppFWURLEncodedFormContentType `json:"appfwurlencodedformcontenttype"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Types) == 0 {
		return nil, fmt.Errorf("appfwurlencodedformcontenttype not found")
	}

	return &result.Types[0], nil
}

func (s *AppFWService) CountAppFWURLEncodedFormContentType() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWURLEncodedFormContentTypeURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Types []struct {
			Count float64 `json:"__count"`
		} `json:"appfwurlencodedformcontenttype"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Types) > 0 {
		return result.Types[0].Count, nil
	}

	return 0, nil
}

// appfwwsdl
// Configuration for WSDL file resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwwsdl
func (s *AppFWService) DeleteAppFWWSDL(name string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWWSDLURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWWSDL() ([]models.AppFWWSDL, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWWSDLURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		WSDLs []models.AppFWWSDL `json:"appfwwsdl"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.WSDLs, nil
}

func (s *AppFWService) GetAppFWWSDL(name string) (*models.AppFWWSDL, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWWSDLURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		WSDLs []models.AppFWWSDL `json:"appfwwsdl"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.WSDLs) == 0 {
		return nil, fmt.Errorf("appfwwsdl not found")
	}

	return &result.WSDLs[0], nil
}

func (s *AppFWService) ImportAppFWWSDL(resource models.AppFWWSDL) error {
	payload := map[string]interface{}{
		"appfwwsdl": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWWSDLURL+"?action=Import", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// appfwxmlcontenttype
// Configuration for XML Content type resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwxmlcontenttype
func (s *AppFWService) AddAppFWXMLContentType(resource models.AppFWXMLContentType) error {
	payload := map[string]interface{}{
		"appfwxmlcontenttype": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWXMLContentTypeURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) DeleteAppFWXMLContentType(name string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWXMLContentTypeURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWXMLContentType() ([]models.AppFWXMLContentType, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWXMLContentTypeURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Types []models.AppFWXMLContentType `json:"appfwxmlcontenttype"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Types, nil
}

func (s *AppFWService) GetAppFWXMLContentType(name string) (*models.AppFWXMLContentType, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWXMLContentTypeURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Types []models.AppFWXMLContentType `json:"appfwxmlcontenttype"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Types) == 0 {
		return nil, fmt.Errorf("appfwxmlcontenttype not found")
	}

	return &result.Types[0], nil
}

func (s *AppFWService) CountAppFWXMLContentType() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWXMLContentTypeURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Types []struct {
			Count float64 `json:"__count"`
		} `json:"appfwxmlcontenttype"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Types) > 0 {
		return result.Types[0].Count, nil
	}

	return 0, nil
}

// appfwxmlerrorpage
// Configuration for xml error page resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwxmlerrorpage
func (s *AppFWService) DeleteAppFWXMLErrorPage(name string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWXMLErrorPageURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWXMLErrorPage() ([]models.AppFWXMLErrorPage, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWXMLErrorPageURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Pages []models.AppFWXMLErrorPage `json:"appfwxmlerrorpage"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Pages, nil
}

func (s *AppFWService) GetAppFWXMLErrorPage(name string) (*models.AppFWXMLErrorPage, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWXMLErrorPageURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Pages []models.AppFWXMLErrorPage `json:"appfwxmlerrorpage"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Pages) == 0 {
		return nil, fmt.Errorf("appfwxmlerrorpage not found")
	}

	return &result.Pages[0], nil
}

func (s *AppFWService) ImportAppFWXMLErrorPage(resource models.AppFWXMLErrorPage) error {
	payload := map[string]interface{}{
		"appfwxmlerrorpage": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWXMLErrorPageURL+"?action=Import", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) ChangeAppFWXMLErrorPage(resource models.AppFWXMLErrorPage) error {
	payload := map[string]interface{}{
		"appfwxmlerrorpage": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWXMLErrorPageURL+"?action=change", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// appfwxmlschema
// Configuration for XML schema resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwxmlschema
func (s *AppFWService) DeleteAppFWXMLSchema(name string) error {
	reqURL := fmt.Sprintf("%s/%s", appFWXMLSchemaURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppFWService) GetAllAppFWXMLSchema() ([]models.AppFWXMLSchema, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFWXMLSchemaURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Schemas []models.AppFWXMLSchema `json:"appfwxmlschema"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Schemas, nil
}

func (s *AppFWService) GetAppFWXMLSchema(name string) (*models.AppFWXMLSchema, error) {
	reqURL := fmt.Sprintf("%s/%s", appFWXMLSchemaURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Schemas []models.AppFWXMLSchema `json:"appfwxmlschema"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Schemas) == 0 {
		return nil, fmt.Errorf("appfwxmlschema not found")
	}

	return &result.Schemas[0], nil
}

func (s *AppFWService) ImportAppFWXMLSchema(resource models.AppFWXMLSchema) error {
	payload := map[string]interface{}{
		"appfwxmlschema": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appFWXMLSchemaURL+"?action=Import", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}
