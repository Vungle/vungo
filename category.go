package openrtb

// Category type represents the IAB standard content category types.
type Category string

// Supported categories
const (
	CATEGORY_ARTS_ENTERTAINMENT   Category = "IAB1"
	CATEGORY_AIDS_HIV             Category = "IAB7-3"
	CATEGORY_ALTERNATIVE_MEDICINE Category = "IAB7-5"
	CATEGORY_INCEST_ABUSE_SUPPORT Category = "IAB7-28"
	CATEGORY_INCONTINENCE         Category = "IAB7-29"
	CATEGORY_INFERTILITY          Category = "IAB7-30"
	CATEGORY_SEXUALITY            Category = "IAB7-39"
	CATEGORY_SMOKING_CESSATION    Category = "IAB7-41"
	CATEGORY_SUBSTANCE_ABUSE      Category = "IAB7-42"
	CATEGORY_WEIGHT_LOSS          Category = "IAB7-44"
	CATEGORY_COCKTAILS_BEER       Category = "IAB8-5"
	CATEGORY_WINE                 Category = "IAB8-18"
	CATEGORY_CIGARS               Category = "IAB9-9"
	CATEGORY_DATING               Category = "IAB14-1"
	CATEGORY_GAY_LIFE             Category = "IAB14-3"
	CATEGORY_MARRIAGE             Category = "IAB14-4"
	CATEGORY_BODY_ART             Category = "IAB18-2"
	CATEGORY_RELIGION             Category = "IAB23"
	CATEGORY_NON_STANDARD_CONTENT Category = "IAB25"
	CATEGORY_ILLEGAL_CONTENT      Category = "IAB26"
)
