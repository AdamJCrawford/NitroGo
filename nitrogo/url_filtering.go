package nitrogo

const (
	urlFilteringCategoriesURL     = "/nitro/v1/config/urlfilteringcategories"
	urlFilteringCategorizationURL = "/nitro/v1/config/urlfilteringcategorization"
	urlFilteringCategoryGroupsURL = "/nitro/v1/config/urlfilteringcategorygroups"
	urlFilteringParameterURL      = "/nitro/v1/config/urlfilteringparameter"
)

// URL Filtering feature is used control access to webpages based on category.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/urlfiltering/urlfiltering
type URLFilteringService struct {
	client *Client
}

// urlfilteringcategories
// Configuration for Categories resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/urlfiltering/urlfilteringcategories
func (s *URLFilteringService) GetAllURLFilteringCategories() {}

// urlfilteringcategorization
// Configuration for Categorization resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/urlfiltering/urlfilteringcategorization
func (s *URLFilteringService) AddURLFilteringCategorization()    {}
func (s *URLFilteringService) ClearURLFilteringCategorization()  {}
func (s *URLFilteringService) GetAllURLFilteringCategorization() {}
func (s *URLFilteringService) CountURLFilteringCategorization()  {}

// urlfilteringcategorygroups
// Configuration for Category Groups resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/urlfiltering/urlfilteringcategorygroups
func (s *URLFilteringService) GetAllURLFilteringCategoryGroups() {}

// urlfilteringparameter
// Configuration for URLFILTERING paramter resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/urlfiltering/urlfilteringparameter
func (s *URLFilteringService) UpdateURLFilteringParameter() {}
func (s *URLFilteringService) UnsetURLFilteringParameter()  {}
func (s *URLFilteringService) GetAllURLFilteringParameter() {}
