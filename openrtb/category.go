package openrtb

// Category type represents the IAB standard content category types.
type Category string

// Possible values of IAB content categories.
const (
	CategoryArtsEntertainment              Category = "IAB1"
	CategoryHobbiesAndInterests            Category = "IAB9"
	CategoryAIDSHIV                        Category = "IAB7-3"
	CategoryAlternativeMedicine            Category = "IAB7-5"
	CategoryIncestAbuseSupport             Category = "IAB7-28"
	CategoryIncontinence                   Category = "IAB7-29"
	CategoryInfertility                    Category = "IAB7-30"
	CategorySexuality                      Category = "IAB7-39"
	CategorySmokingCessation               Category = "IAB7-41"
	CategorySubstanceAbuse                 Category = "IAB7-42"
	CategoryWeightLoss                     Category = "IAB7-44"
	CategoryCocktailsBeer                  Category = "IAB8-5"
	CategoryWine                           Category = "IAB8-18"
	CategoryCigars                         Category = "IAB9-9"
	CategoryDating                         Category = "IAB14-1"
	CategoryGayLife                        Category = "IAB14-3"
	CategoryMarriage                       Category = "IAB14-4"
	CategoryBodyArt                        Category = "IAB18-2"
	CategoryReligion                       Category = "IAB23"
	CategoryNonStandardContent             Category = "IAB25"
	CategoryUnmoderatedUGC                 Category = "IAB25-1"
	CategoryExtremeGraphicExplicitViolence Category = "IAB25-2"
	CategoryPornography                    Category = "IAB25-3"
	CategoryProfaneContent                 Category = "IAB25-4"
	CategoryHateContent                    Category = "IAB25-5"
	CategoryUnderConstruction              Category = "IAB25-6"
	CategoryIncentivized                   Category = "IAB25-7"
	CategoryIllegalContent                 Category = "IAB26"
)
