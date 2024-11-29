package nitrogo

type URLFilteringService struct {
	client *Client
}

// urlfilteringcategories
func (s *URLFilteringService) GetAllURLFilteringCategories() {}

// urlfilteringcategorization
func (s *URLFilteringService) AddURLFilteringCategorization()    {}
func (s *URLFilteringService) ClearURLFilteringCategorization()  {}
func (s *URLFilteringService) GetAllURLFilteringCategorization() {}
func (s *URLFilteringService) CountURLFilteringCategorization()  {}

// urlfilteringcategorygroups
func (s *URLFilteringService) GetAllURLFilteringCategoryGroups() {}

// urlfilteringparameter
func (s *URLFilteringService) UpdateURLFilteringParameter() {}
func (s *URLFilteringService) UnsetURLFilteringParameter()  {}
func (s *URLFilteringService) GetAllURLFilteringParameter() {}
