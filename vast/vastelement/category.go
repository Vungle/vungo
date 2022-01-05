package vastelement

// Category type is used in creative separation and for compliance in certain programs,
// a category field is needed to categorize the ad’s content. Vast 4.0.
type Category struct {
	Authority string      `xml:"authority,attr"`
	Content   TrimmedData `xml:",chardata"`
}

// Validate method validate Category vast element
func (category *Category) Validate(version Version) error {
	if len(category.Authority) == 0 {
		return ValidationError{Errs: []error{ErrCategoryMissingAuthority}}
	}
	return nil
}
