package openrtb

// Category type represents the IAB standard content category types.
type Category string

// Possible values of IAB content categories.
const (
	CategoryArtsEntertainment   Category = "IAB1"
	CategoryAIDSHIV             Category = "IAB7-3"
	CategoryAlternativeMedicine Category = "IAB7-5"
	CategoryIncestAbuseSupport  Category = "IAB7-28"
	CategoryIncontinence        Category = "IAB7-29"
	CategoryInfertility         Category = "IAB7-30"
	CategorySexuality           Category = "IAB7-39"
	CategorySmokingCessation    Category = "IAB7-41"
	CategorySubstanceAbuse      Category = "IAB7-42"
	CategoryWeightLoss          Category = "IAB7-44"
	CategoryCocktailsBeer       Category = "IAB8-5"
	CategoryWine                Category = "IAB8-18"
	CategoryCigars              Category = "IAB9-9"
	CategoryDating              Category = "IAB14-1"
	CategoryGayLife             Category = "IAB14-3"
	CategoryMarriage            Category = "IAB14-4"
	CategoryBodyArt             Category = "IAB18-2"
	CategoryReligion            Category = "IAB23"
	CategoryNonStandardContent  Category = "IAB25"
	CategoryIllegalContent      Category = "IAB26"
)
