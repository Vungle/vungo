package openrtb

import "fmt"

// Category type represents the IAB standard content category types.
type Category string

// Possible values of IAB content categories.
const (
	CategoryArtsAndEntertainment Category = "IAB1"
	CategoryBooksAndLiterature   Category = "IAB1-1"
	CategoryCelebrityFanGossip   Category = "IAB1-2"
	CategoryFineArt              Category = "IAB1-3"
	CategoryHumor                Category = "IAB1-4"
	CategoryMovies               Category = "IAB1-5"
	CategoryMusic                Category = "IAB1-6"
	CategoryTelevision           Category = "IAB1-7"

	CategoryAutomotive           Category = "IAB2"
	CategoryAutoParts            Category = "IAB2-1"
	CategoryAutoRepair           Category = "IAB2-2"
	CategoryBuyingSellingCars    Category = "IAB2-3"
	CategoryCarCulture           Category = "IAB2-4"
	CategoryCertifiedPreOwned    Category = "IAB2-5"
	CategoryConvertible          Category = "IAB2-6"
	CategoryCoupe                Category = "IAB2-7"
	CategoryCrossover            Category = "IAB2-8"
	CategoryDiesel               Category = "IAB2-9"
	CategoryElectricVehicle      Category = "IAB2-10"
	CategoryHatchback            Category = "IAB2-11"
	CategoryHybrid               Category = "IAB2-12"
	CategoryLuxury               Category = "IAB2-13"
	CategoryMinivan              Category = "IAB2-14"
	CategoryMororcycles          Category = "IAB2-15"
	CategoryOffRoadVehicles      Category = "IAB2-16"
	CategoryPerformanceVehicles  Category = "IAB2-17"
	CategoryPickup               Category = "IAB2-18"
	CategoryRoadSideAssistance   Category = "IAB2-19"
	CategorySedan                Category = "IAB2-20"
	CategoryTrucksAndAccessories Category = "IAB2-21"
	CategoryVintageCars          Category = "IAB2-22"
	CategoryWagon                Category = "IAB2-23"

	CategoryBusiness          Category = "IAB3"
	CategoryAdvertising       Category = "IAB3-1"
	CategoryAgriculture       Category = "IAB3-2"
	CategoryBiotechBiomedical Category = "IAB3-3"
	CategoryBusinessSoftware  Category = "IAB3-4"
	CategoryConstruction      Category = "IAB3-5"
	CategoryForestry          Category = "IAB3-6"
	CategoryGovernment        Category = "IAB3-7"
	CategoryGreenSolutions    Category = "IAB3-8"
	CategoryHumanResources    Category = "IAB3-9"
	CategoryLogistics         Category = "IAB3-10"
	CategoryMarketing         Category = "IAB3-11"
	CategoryMetals            Category = "IAB3-12"

	CategoryCareers             Category = "IAB4"
	CategoryCareerPlanning      Category = "IAB4-1"
	CategoryCollege             Category = "IAB4-2"
	CategoryFinancialAid        Category = "IAB4-3"
	CategoryJobFairs            Category = "IAB4-4"
	CategoryJobSearch           Category = "IAB4-5"
	CategoryResumeWritingAdvice Category = "IAB4-6"
	CategoryNursing             Category = "IAB4-7"
	CategoryScholarships        Category = "IAB4-8"
	CategoryTelecommuting       Category = "IAB4-9"
	CategoryUSMilitary          Category = "IAB4-10"
	CategoryCareerAdvice        Category = "IAB4-11"

	CategoryEducation              Category = "IAB5"
	Category712Education           Category = "IAB5-1"
	CategoryAdultEducation         Category = "IAB5-2"
	CategoryArtHistory             Category = "IAB5-3"
	CategoryColledgeAdministration Category = "IAB5-4"
	CategoryCollegeLife            Category = "IAB5-5"
	CategoryDistanceLearning       Category = "IAB5-6"
	CategoryEnglishAsA2ndLanguage  Category = "IAB5-7"
	CategoryLanguageLearning       Category = "IAB5-8"
	CategoryGraduateSchool         Category = "IAB5-9"
	CategoryHomeschooling          Category = "IAB5-10"
	CategoryHomeworkStudyTips      Category = "IAB5-11"
	CategoryK6Educators            Category = "IAB5-12"
	CategoryPrivateSchool          Category = "IAB5-13"
	CategorySpecialEducation       Category = "IAB5-14"
	CategoryStudyingBusiness       Category = "IAB5-15"

	CategoryFamilyAndParenting Category = "IAB6"
	CategoryAdoption           Category = "IAB6-1"
	CategoryBabiesAndToddlers  Category = "IAB6-2"
	CategoryDaycarePreSchool   Category = "IAB6-3"
	CategoryFamilyInternet     Category = "IAB6-4"
	CategoryParentingK6Kids    Category = "IAB6-5"
	CategoryParentingTeens     Category = "IAB6-6"
	CategoryPregnancy          Category = "IAB6-7"
	CategorySpecialNeedsKids   Category = "IAB6-8"
	CategoryEldercare          Category = "IAB6-9"

	CategoryHealthAndFitness       Category = "IAB7"
	CategoryExercise               Category = "IAB7-1"
	CategoryADD                    Category = "IAB7-2"
	CategoryAIDSHIV                Category = "IAB7-3"
	CategoryAllergies              Category = "IAB7-4"
	CategoryAlternativeMedicine    Category = "IAB7-5"
	CategoryArthritis              Category = "IAB7-6"
	CategoryAsthma                 Category = "IAB7-7"
	CategoryAutismPDD              Category = "IAB7-8"
	CategoryBipolarDisorder        Category = "IAB7-9"
	CategoryBrainTumor             Category = "IAB7-10"
	CategoryCancer                 Category = "IAB7-11"
	CategoryCholesterol            Category = "IAB7-12"
	CategoryChronicFatigueSyndrome Category = "IAB7-13"
	CategoryChronicPain            Category = "IAB7-14"
	CategoryColdAndFlu             Category = "IAB7-15"
	CategoryDeafness               Category = "IAB7-16"
	CategoryDentalCare             Category = "IAB7-17"
	CategoryDepression             Category = "IAB7-18"
	CategoryDermatology            Category = "IAB7-19"
	CategoryDiabetes               Category = "IAB7-20"
	CategoryEpilepsy               Category = "IAB7-21"
	CategoryGERDAcidReflux         Category = "IAB7-22"
	CategoryHeadachesMigraines     Category = "IAB7-23"
	CategoryHeartDisease           Category = "IAB7-24"
	CategoryHerbsforHealth         Category = "IAB7-25"
	CategoryHolisticHealing        Category = "IAB7-26"
	CategoryIBSCrohnsDisease       Category = "IAB7-27"
	CategoryIncestAbuseSupport     Category = "IAB7-28"
	CategoryIncontinence           Category = "IAB7-29"
	CategoryInfertility            Category = "IAB7-30"
	CategoryMensHealth             Category = "IAB7-31"
	CategoryNutrition              Category = "IAB7-32"
	CategoryOrthopedics            Category = "IAB7-33"
	CategoryPanicAnxietyDisorders  Category = "IAB7-34"
	CategoryPediatrics             Category = "IAB7-35"
	CategoryPhysicalTherapy        Category = "IAB7-36"
	CategoryPsychologyPsychiatry   Category = "IAB7-37"
	CategorySeniorHealth           Category = "IAB7-38"
	CategorySexuality              Category = "IAB7-39"
	CategorySleepDisorders         Category = "IAB7-40"
	CategorySmokingCessation       Category = "IAB7-41"
	CategorySubstanceAbuse         Category = "IAB7-42"
	CategoryThyroidDisease         Category = "IAB7-43"
	CategoryWeightLoss             Category = "IAB7-44"
	CategoryWomensHealth           Category = "IAB7-45"

	CategoryFoodAndDrink         Category = "IAB8"
	CategoryAmericanCuisine      Category = "IAB8-1"
	CategoryBarbecuesAndGrilling Category = "IAB8-2"
	CategoryCajunCreole          Category = "IAB8-3"
	CategoryChineseCuisine       Category = "IAB8-4"
	CategoryCocktailsBeer        Category = "IAB8-5"
	CategoryCoffeeTea            Category = "IAB8-6"
	CategoryCuisineSpecific      Category = "IAB8-7"
	CategoryDessertsAndBaking    Category = "IAB8-8"
	CategoryDiningOut            Category = "IAB8-9"
	CategoryFoodAllergies        Category = "IAB8-10"
	CategoryFrenchCuisine        Category = "IAB8-11"
	CategoryHealthLowFatCooking  Category = "IAB8-12"
	CategoryItalianCuisine       Category = "IAB8-13"
	CategoryJapaneseCuisine      Category = "IAB8-14"
	CategoryMexicanCuisine       Category = "IAB8-15"
	CategoryVegan                Category = "IAB8-16"
	CategoryVegetarian           Category = "IAB8-17"
	CategoryWine                 Category = "IAB8-18"

	CategoryHobbiesAndInterests   Category = "IAB9"
	CategoryArtTechnology         Category = "IAB9-1"
	CategoryArtsAndCrafts         Category = "IAB9-2"
	CategoryBeadwork              Category = "IAB9-3"
	CategoryBirdwatching          Category = "IAB9-4"
	CategoryBoardGamesPuzzles     Category = "IAB9-5"
	CategoryCandleAndSoapMaking   Category = "IAB9-6"
	CategoryCardGames             Category = "IAB9-7"
	CategoryChess                 Category = "IAB9-8"
	CategoryCigars                Category = "IAB9-9"
	CategoryCollecting            Category = "IAB9-10"
	CategoryComicBooks            Category = "IAB9-11"
	CategoryDrawingSketching      Category = "IAB9-12"
	CategoryFreelanceWriting      Category = "IAB9-13"
	CategoryGenealogy             Category = "IAB9-14"
	CategoryGettingPublished      Category = "IAB9-15"
	CategoryGuitar                Category = "IAB9-16"
	CategoryHomeRecording         Category = "IAB9-17"
	CategoryInvestorsAndPatents   Category = "IAB9-18"
	CategoryJewelryMaking         Category = "IAB9-19"
	CategoryMagicAndIllusion      Category = "IAB9-20"
	CategoryNeedlework            Category = "IAB9-21"
	CategoryPainting              Category = "IAB9-22"
	CategoryPhotography           Category = "IAB9-23"
	CategoryRadio                 Category = "IAB9-24"
	CategoryRoleplayingGames      Category = "IAB9-25"
	CategorySciFiAndFantasy       Category = "IAB9-26"
	CategoryScrapbooking          Category = "IAB9-27"
	CategoryScreenwriting         Category = "IAB9-28"
	CategoryStampsAndCoins        Category = "IAB9-29"
	CategoryVideoAndComputerGames Category = "IAB9-30"
	CategoryWoodworking           Category = "IAB9-31"

	CategoryHomeAndGarden             Category = "IAB10"
	CategoryAppliances                Category = "IAB10-1"
	CategoryEntertaining              Category = "IAB10-2"
	CategoryEnvironmentalSafety       Category = "IAB10-3"
	CategoryGardening                 Category = "IAB10-4"
	CategoryHomeRepair                Category = "IAB10-5"
	CategoryHomeTheater               Category = "IAB10-6"
	CategoryInteriorDecorating        Category = "IAB10-7"
	CategoryLandscaping               Category = "IAB10-8"
	CategoryRemodelingAndConstruction Category = "IAB10-9"

	CategoryLawGovernmentAndPolitics Category = "IAB11"
	CategoryImmigration              Category = "IAB11-1"
	CategoryLegalIssues              Category = "IAB11-2"
	CategoryUSGovernmentResources    Category = "IAB11-3"
	CategoryPolitics                 Category = "IAB11-4"
	CategoryCommentary               Category = "IAB11-5"
	CategoryNews                     Category = "IAB12"
	CategoryInternationalNews        Category = "IAB12-1"
	CategoryNationalNews             Category = "IAB12-2"
	CategoryLocalNews                Category = "IAB12-3"

	CategoryPersonalFinance    Category = "IAB13"
	CategoryBeginningInvesting Category = "IAB13-1"
	CategoryCreditDebtAndLoans Category = "IAB13-2"
	CategoryFinancialNews      Category = "IAB13-3"
	CategoryFinancialPlanning  Category = "IAB13-4"
	CategoryHedgeFund          Category = "IAB13-5"
	CategoryInsurance          Category = "IAB13-6"
	CategoryInvesting          Category = "IAB13-7"
	CategoryMutualFunds        Category = "IAB13-8"
	CategoryOptions            Category = "IAB13-9"
	CategoryRetirementPlanning Category = "IAB13-10"
	CategoryStocks             Category = "IAB13-11"
	CategoryTaxPlanning        Category = "IAB13-12"

	CategorySociety        Category = "IAB14"
	CategoryDating         Category = "IAB14-1"
	CategoryDivorceSupport Category = "IAB14-2"
	CategoryGayLife        Category = "IAB14-3"
	CategoryMarriage       Category = "IAB14-4"
	CategorySeniorLiving   Category = "IAB14-5"
	CategoryTeens          Category = "IAB14-6"
	CategoryWeddings       Category = "IAB14-7"
	CategoryEthnicSpecific Category = "IAB14-8"

	CategoryScience             Category = "IAB15"
	CategoryAstrology           Category = "IAB15-1"
	CategoryBiology             Category = "IAB15-2"
	CategoryChemistry           Category = "IAB15-3"
	CategoryGeology             Category = "IAB15-4"
	CategoryParanormalPhenomena Category = "IAB15-5"
	CategoryPhysics             Category = "IAB15-6"
	CategorySpaceAstronomy      Category = "IAB15-7"
	CategoryGeography           Category = "IAB15-8"
	CategoryBotany              Category = "IAB15-9"
	CategoryWeather             Category = "IAB15-10"

	CategoryPets               Category = "IAB16"
	CategoryAquariums          Category = "IAB16-1"
	CategoryBirds              Category = "IAB16-2"
	CategoryCats               Category = "IAB16-3"
	CategoryDogs               Category = "IAB16-4"
	CategoryLargeAnimals       Category = "IAB16-5"
	CategoryReptiles           Category = "IAB16-6"
	CategoryVeterinaryMedicine Category = "IAB16-7"

	CategorySports              Category = "IAB17"
	CategoryAutoRacing          Category = "IAB17-1"
	CategoryBaseball            Category = "IAB17-2"
	CategoryBicycling           Category = "IAB17-3"
	CategoryBodybuilding        Category = "IAB17-4"
	CategoryBoxing              Category = "IAB17-5"
	CategoryCanoeingKayaking    Category = "IAB17-6"
	CategoryCheerleading        Category = "IAB17-7"
	CategoryClimbing            Category = "IAB17-8"
	CategoryCricket             Category = "IAB17-9"
	CategoryFigureSkating       Category = "IAB17-10"
	CategoryFlyFishing          Category = "IAB17-11"
	CategoryFootball            Category = "IAB17-12"
	CategoryFreshwaterFishing   Category = "IAB17-13"
	CategoryGameAndFish         Category = "IAB17-14"
	CategoryGolf                Category = "IAB17-15"
	CategoryHorseRacing         Category = "IAB17-16"
	CategoryHorses              Category = "IAB17-17"
	CategoryHuntingShooting     Category = "IAB17-18"
	CategoryInlineSkating       Category = "IAB17-19"
	CategoryMartialArts         Category = "IAB17-20"
	CategoryMountainBiking      Category = "IAB17-21"
	CategoryNASCARRacing        Category = "IAB17-22"
	CategoryOlympics            Category = "IAB17-23"
	CategoryPaintball           Category = "IAB17-24"
	CategoryPowerAndMotorcycles Category = "IAB17-25"
	CategoryProBasketball       Category = "IAB17-26"
	CategoryProIceHockey        Category = "IAB17-27"
	CategoryRodeo               Category = "IAB17-28"
	CategoryRugby               Category = "IAB17-29"
	CategoryRunningJogging      Category = "IAB17-30"
	CategorySailing             Category = "IAB17-31"
	CategorySaltwaterFishing    Category = "IAB17-32"
	CategoryScubaDiving         Category = "IAB17-33"
	CategorySkateboarding       Category = "IAB17-34"
	CategorySkiing              Category = "IAB17-35"
	CategorySnowboarding        Category = "IAB17-36"
	CategorySurfingBodyboarding Category = "IAB17-37"
	CategorySwimming            Category = "IAB17-38"
	CategoryTableTennisPingPong Category = "IAB17-39"
	CategoryTennis              Category = "IAB17-40"
	CategoryVolleyball          Category = "IAB17-41"
	CategoryWalking             Category = "IAB17-42"
	CategoryWaterskiWakeboard   Category = "IAB17-43"
	CategoryWorldSoccer         Category = "IAB17-44"

	CategoryStyleAndFashion Category = "IAB18"
	CategoryBeauty          Category = "IAB18-1"
	CategoryBodyArt         Category = "IAB18-2"
	CategoryFashion         Category = "IAB18-3"
	CategoryJewelry         Category = "IAB18-4"
	CategoryClothing        Category = "IAB18-5"
	CategoryAccessories     Category = "IAB18-6"

	CategoryTechnologyAndComputing Category = "IAB19"
	Category3DGraphics             Category = "IAB19-1"
	CategoryAnimation              Category = "IAB19-2"
	CategoryAntivirusSoftware      Category = "IAB19-3"
	CategoryCCPlusPlus             Category = "IAB19-4"
	CategoryCamerasAndCamcorders   Category = "IAB19-5"
	CategoryCellPhones             Category = "IAB19-6"
	CategoryComputerCertification  Category = "IAB19-7"
	CategoryComputerNetworking     Category = "IAB19-8"
	CategoryComputerPeripherals    Category = "IAB19-9"
	CategoryComputerReviews        Category = "IAB19-10"
	CategoryDataCenters            Category = "IAB19-11"
	CategoryDatabases              Category = "IAB19-12"
	CategoryDesktopPublishing      Category = "IAB19-13"
	CategoryDesktopVideo           Category = "IAB19-14"
	CategoryEmail                  Category = "IAB19-15"
	CategoryGraphicsSoftware       Category = "IAB19-16"
	CategoryHomeVideoDVD           Category = "IAB19-17"
	CategoryInternetTechnology     Category = "IAB19-18"
	CategoryJava                   Category = "IAB19-19"
	CategoryJavaScript             Category = "IAB19-20"
	CategoryMacSupport             Category = "IAB19-21"
	CategoryMP3MIDI                Category = "IAB19-22"
	CategoryNetConferencing        Category = "IAB19-23"
	CategoryNetforBeginners        Category = "IAB19-24"
	CategoryNetworkSecurity        Category = "IAB19-25"
	CategoryPalmtopsPDAs           Category = "IAB19-26"
	CategoryPCSupport              Category = "IAB19-27"
	CategoryPortable               Category = "IAB19-28"
	CategoryEntertainment          Category = "IAB19-29"
	CategorySharewareFreeware      Category = "IAB19-30"
	CategoryUnix                   Category = "IAB19-31"
	CategoryVisualBasic            Category = "IAB19-32"
	CategoryWebClipArt             Category = "IAB19-33"
	CategoryWebDesignHTML          Category = "IAB19-34"
	CategoryWebSearch              Category = "IAB19-35"
	CategoryWindows                Category = "IAB19-36"

	CategoryTravel                  Category = "IAB20"
	CategoryAdventureTravel         Category = "IAB20-1"
	CategoryAfrica                  Category = "IAB20-2"
	CategoryAirTravel               Category = "IAB20-3"
	CategoryAustraliaAndNewZealand  Category = "IAB20-4"
	CategoryBedAndBreakfasts        Category = "IAB20-5"
	CategoryBudgetTravel            Category = "IAB20-6"
	CategoryBusinessTravel          Category = "IAB20-7"
	CategoryByUSLocale              Category = "IAB20-8"
	CategoryCamping                 Category = "IAB20-9"
	CategoryCanada                  Category = "IAB20-10"
	CategoryCaribbean               Category = "IAB20-11"
	CategoryCruises                 Category = "IAB20-12"
	CategoryEasternEurope           Category = "IAB20-13"
	CategoryEurope                  Category = "IAB20-14"
	CategoryFrance                  Category = "IAB20-15"
	CategoryGreece                  Category = "IAB20-16"
	CategoryHoneymoonsGetaways      Category = "IAB20-17"
	CategoryHotels                  Category = "IAB20-18"
	CategoryItaly                   Category = "IAB20-19"
	CategoryJapan                   Category = "IAB20-20"
	CategoryMexicoAndCentralAmerica Category = "IAB20-21"
	CategoryNationalParks           Category = "IAB20-22"
	CategorySouthAmerica            Category = "IAB20-23"
	CategorySpas                    Category = "IAB20-24"
	CategoryThemeParks              Category = "IAB20-25"
	CategoryTravelingwithKids       Category = "IAB20-26"
	CategoryUnitedKingdom           Category = "IAB20-27"

	CategoryRealEstate         Category = "IAB21"
	CategoryApartments         Category = "IAB21-1"
	CategoryArchitects         Category = "IAB21-2"
	CategoryBuyingSellingHomes Category = "IAB21-3"

	CategoryShopping            Category = "IAB22"
	CategoryContestsAndFreebies Category = "IAB22-1"
	CategoryCouponing           Category = "IAB22-2"
	CategoryComparison          Category = "IAB22-3"
	CategoryEngines             Category = "IAB22-4"

	CategoryReligionAndSpirituality Category = "IAB23"
	CategoryAlternativeReligions    Category = "IAB23-1"
	CategoryAtheismAgnosticism      Category = "IAB23-2"
	CategoryBuddhism                Category = "IAB23-3"
	CategoryCatholicism             Category = "IAB23-4"
	CategoryChristianity            Category = "IAB23-5"
	CategoryHinduism                Category = "IAB23-6"
	CategoryIslam                   Category = "IAB23-7"
	CategoryJudaism                 Category = "IAB23-8"
	CategoryLatterDaySaints         Category = "IAB23-9"
	CategoryPaganWiccan             Category = "IAB23-10"

	CategoryUncategorized Category = "IAB24"

	CategoryNonStandardContent             Category = "IAB25"
	CategoryUnmoderatedUGC                 Category = "IAB25-1"
	CategoryExtremeGraphicExplicitViolence Category = "IAB25-2"
	CategoryPornography                    Category = "IAB25-3"
	CategoryProfaneContent                 Category = "IAB25-4"
	CategoryHateContent                    Category = "IAB25-5"
	CategoryUnderConstruction              Category = "IAB25-6"
	CategoryIncentivized                   Category = "IAB25-7"

	CategoryIllegalContent        Category = "IAB26"
	CategoryIllegalContent2       Category = "IAB26-1"
	CategoryWarez                 Category = "IAB26-2"
	CategorySpywareMalware        Category = "IAB26-3"
	CategoryCopyrightInfringement Category = "IAB26-4"
)

var categoryNameToValue = map[string]Category{
	string(CategoryArtsAndEntertainment): CategoryArtsAndEntertainment,
	string(CategoryBooksAndLiterature):   CategoryBooksAndLiterature,
	string(CategoryCelebrityFanGossip):   CategoryCelebrityFanGossip,
	string(CategoryFineArt):              CategoryFineArt,
	string(CategoryHumor):                CategoryHumor,
	string(CategoryMovies):               CategoryMovies,
	string(CategoryMusic):                CategoryMusic,
	string(CategoryTelevision):           CategoryTelevision,

	string(CategoryAutomotive):           CategoryAutomotive,
	string(CategoryAutoParts):            CategoryAutoParts,
	string(CategoryAutoRepair):           CategoryAutoRepair,
	string(CategoryBuyingSellingCars):    CategoryBuyingSellingCars,
	string(CategoryCarCulture):           CategoryCarCulture,
	string(CategoryCertifiedPreOwned):    CategoryCertifiedPreOwned,
	string(CategoryConvertible):          CategoryConvertible,
	string(CategoryCoupe):                CategoryCoupe,
	string(CategoryCrossover):            CategoryCrossover,
	string(CategoryDiesel):               CategoryDiesel,
	string(CategoryElectricVehicle):      CategoryElectricVehicle,
	string(CategoryHatchback):            CategoryHatchback,
	string(CategoryHybrid):               CategoryHybrid,
	string(CategoryLuxury):               CategoryLuxury,
	string(CategoryMinivan):              CategoryMinivan,
	string(CategoryMororcycles):          CategoryMororcycles,
	string(CategoryOffRoadVehicles):      CategoryOffRoadVehicles,
	string(CategoryPerformanceVehicles):  CategoryPerformanceVehicles,
	string(CategoryPickup):               CategoryPickup,
	string(CategoryRoadSideAssistance):   CategoryRoadSideAssistance,
	string(CategorySedan):                CategorySedan,
	string(CategoryTrucksAndAccessories): CategoryTrucksAndAccessories,
	string(CategoryVintageCars):          CategoryVintageCars,
	string(CategoryWagon):                CategoryWagon,

	string(CategoryBusiness):          CategoryBusiness,
	string(CategoryAdvertising):       CategoryAdvertising,
	string(CategoryAgriculture):       CategoryAgriculture,
	string(CategoryBiotechBiomedical): CategoryBiotechBiomedical,
	string(CategoryBusinessSoftware):  CategoryBusinessSoftware,
	string(CategoryConstruction):      CategoryConstruction,
	string(CategoryForestry):          CategoryForestry,
	string(CategoryGovernment):        CategoryGovernment,
	string(CategoryGreenSolutions):    CategoryGreenSolutions,
	string(CategoryHumanResources):    CategoryHumanResources,
	string(CategoryLogistics):         CategoryLogistics,
	string(CategoryMarketing):         CategoryMarketing,
	string(CategoryMetals):            CategoryMetals,

	string(CategoryCareers):             CategoryCareers,
	string(CategoryCareerPlanning):      CategoryCareerPlanning,
	string(CategoryCollege):             CategoryCollege,
	string(CategoryFinancialAid):        CategoryFinancialAid,
	string(CategoryJobFairs):            CategoryJobFairs,
	string(CategoryJobSearch):           CategoryJobSearch,
	string(CategoryResumeWritingAdvice): CategoryResumeWritingAdvice,
	string(CategoryNursing):             CategoryNursing,
	string(CategoryScholarships):        CategoryScholarships,
	string(CategoryTelecommuting):       CategoryTelecommuting,
	string(CategoryUSMilitary):          CategoryUSMilitary,
	string(CategoryCareerAdvice):        CategoryCareerAdvice,

	string(CategoryEducation):              CategoryEducation,
	string(Category712Education):           Category712Education,
	string(CategoryAdultEducation):         CategoryAdultEducation,
	string(CategoryArtHistory):             CategoryArtHistory,
	string(CategoryColledgeAdministration): CategoryColledgeAdministration,
	string(CategoryCollegeLife):            CategoryCollegeLife,
	string(CategoryDistanceLearning):       CategoryDistanceLearning,
	string(CategoryEnglishAsA2ndLanguage):  CategoryEnglishAsA2ndLanguage,
	string(CategoryLanguageLearning):       CategoryLanguageLearning,
	string(CategoryGraduateSchool):         CategoryGraduateSchool,
	string(CategoryHomeschooling):          CategoryHomeschooling,
	string(CategoryHomeworkStudyTips):      CategoryHomeworkStudyTips,
	string(CategoryK6Educators):            CategoryK6Educators,
	string(CategoryPrivateSchool):          CategoryPrivateSchool,
	string(CategorySpecialEducation):       CategorySpecialEducation,
	string(CategoryStudyingBusiness):       CategoryStudyingBusiness,

	string(CategoryFamilyAndParenting): CategoryFamilyAndParenting,
	string(CategoryAdoption):           CategoryAdoption,
	string(CategoryBabiesAndToddlers):  CategoryBabiesAndToddlers,
	string(CategoryDaycarePreSchool):   CategoryDaycarePreSchool,
	string(CategoryFamilyInternet):     CategoryFamilyInternet,
	string(CategoryParentingK6Kids):    CategoryParentingK6Kids,
	string(CategoryParentingTeens):     CategoryParentingTeens,
	string(CategoryPregnancy):          CategoryPregnancy,
	string(CategorySpecialNeedsKids):   CategorySpecialNeedsKids,
	string(CategoryEldercare):          CategoryEldercare,

	string(CategoryHealthAndFitness):       CategoryHealthAndFitness,
	string(CategoryExercise):               CategoryExercise,
	string(CategoryADD):                    CategoryADD,
	string(CategoryAIDSHIV):                CategoryAIDSHIV,
	string(CategoryAllergies):              CategoryAllergies,
	string(CategoryAlternativeMedicine):    CategoryAlternativeMedicine,
	string(CategoryArthritis):              CategoryArthritis,
	string(CategoryAsthma):                 CategoryAsthma,
	string(CategoryAutismPDD):              CategoryAutismPDD,
	string(CategoryBipolarDisorder):        CategoryBipolarDisorder,
	string(CategoryBrainTumor):             CategoryBrainTumor,
	string(CategoryCancer):                 CategoryCancer,
	string(CategoryCholesterol):            CategoryCholesterol,
	string(CategoryChronicFatigueSyndrome): CategoryChronicFatigueSyndrome,
	string(CategoryChronicPain):            CategoryChronicPain,
	string(CategoryColdAndFlu):             CategoryColdAndFlu,
	string(CategoryDeafness):               CategoryDeafness,
	string(CategoryDentalCare):             CategoryDentalCare,
	string(CategoryDepression):             CategoryDepression,
	string(CategoryDermatology):            CategoryDermatology,
	string(CategoryDiabetes):               CategoryDiabetes,
	string(CategoryEpilepsy):               CategoryEpilepsy,
	string(CategoryGERDAcidReflux):         CategoryGERDAcidReflux,
	string(CategoryHeadachesMigraines):     CategoryHeadachesMigraines,
	string(CategoryHeartDisease):           CategoryHeartDisease,
	string(CategoryHerbsforHealth):         CategoryHerbsforHealth,
	string(CategoryHolisticHealing):        CategoryHolisticHealing,
	string(CategoryIBSCrohnsDisease):       CategoryIBSCrohnsDisease,
	string(CategoryIncestAbuseSupport):     CategoryIncestAbuseSupport,
	string(CategoryIncontinence):           CategoryIncontinence,
	string(CategoryInfertility):            CategoryInfertility,
	string(CategoryMensHealth):             CategoryMensHealth,
	string(CategoryNutrition):              CategoryNutrition,
	string(CategoryOrthopedics):            CategoryOrthopedics,
	string(CategoryPanicAnxietyDisorders):  CategoryPanicAnxietyDisorders,
	string(CategoryPediatrics):             CategoryPediatrics,
	string(CategoryPhysicalTherapy):        CategoryPhysicalTherapy,
	string(CategoryPsychologyPsychiatry):   CategoryPsychologyPsychiatry,
	string(CategorySeniorHealth):           CategorySeniorHealth,
	string(CategorySexuality):              CategorySexuality,
	string(CategorySleepDisorders):         CategorySleepDisorders,
	string(CategorySmokingCessation):       CategorySmokingCessation,
	string(CategorySubstanceAbuse):         CategorySubstanceAbuse,
	string(CategoryThyroidDisease):         CategoryThyroidDisease,
	string(CategoryWeightLoss):             CategoryWeightLoss,
	string(CategoryWomensHealth):           CategoryWomensHealth,

	string(CategoryFoodAndDrink):         CategoryFoodAndDrink,
	string(CategoryAmericanCuisine):      CategoryAmericanCuisine,
	string(CategoryBarbecuesAndGrilling): CategoryBarbecuesAndGrilling,
	string(CategoryCajunCreole):          CategoryCajunCreole,
	string(CategoryChineseCuisine):       CategoryChineseCuisine,
	string(CategoryCocktailsBeer):        CategoryCocktailsBeer,
	string(CategoryCoffeeTea):            CategoryCoffeeTea,
	string(CategoryCuisineSpecific):      CategoryCuisineSpecific,
	string(CategoryDessertsAndBaking):    CategoryDessertsAndBaking,
	string(CategoryDiningOut):            CategoryDiningOut,
	string(CategoryFoodAllergies):        CategoryFoodAllergies,
	string(CategoryFrenchCuisine):        CategoryFrenchCuisine,
	string(CategoryHealthLowFatCooking):  CategoryHealthLowFatCooking,
	string(CategoryItalianCuisine):       CategoryItalianCuisine,
	string(CategoryJapaneseCuisine):      CategoryJapaneseCuisine,
	string(CategoryMexicanCuisine):       CategoryMexicanCuisine,
	string(CategoryVegan):                CategoryVegan,
	string(CategoryVegetarian):           CategoryVegetarian,
	string(CategoryWine):                 CategoryWine,

	string(CategoryHobbiesAndInterests):   CategoryHobbiesAndInterests,
	string(CategoryArtTechnology):         CategoryArtTechnology,
	string(CategoryArtsAndCrafts):         CategoryArtsAndCrafts,
	string(CategoryBeadwork):              CategoryBeadwork,
	string(CategoryBirdwatching):          CategoryBirdwatching,
	string(CategoryBoardGamesPuzzles):     CategoryBoardGamesPuzzles,
	string(CategoryCandleAndSoapMaking):   CategoryCandleAndSoapMaking,
	string(CategoryCardGames):             CategoryCardGames,
	string(CategoryChess):                 CategoryChess,
	string(CategoryCigars):                CategoryCigars,
	string(CategoryCollecting):            CategoryCollecting,
	string(CategoryComicBooks):            CategoryComicBooks,
	string(CategoryDrawingSketching):      CategoryDrawingSketching,
	string(CategoryFreelanceWriting):      CategoryFreelanceWriting,
	string(CategoryGenealogy):             CategoryGenealogy,
	string(CategoryGettingPublished):      CategoryGettingPublished,
	string(CategoryGuitar):                CategoryGuitar,
	string(CategoryHomeRecording):         CategoryHomeRecording,
	string(CategoryInvestorsAndPatents):   CategoryInvestorsAndPatents,
	string(CategoryJewelryMaking):         CategoryJewelryMaking,
	string(CategoryMagicAndIllusion):      CategoryMagicAndIllusion,
	string(CategoryNeedlework):            CategoryNeedlework,
	string(CategoryPainting):              CategoryPainting,
	string(CategoryPhotography):           CategoryPhotography,
	string(CategoryRadio):                 CategoryRadio,
	string(CategoryRoleplayingGames):      CategoryRoleplayingGames,
	string(CategorySciFiAndFantasy):       CategorySciFiAndFantasy,
	string(CategoryScrapbooking):          CategoryScrapbooking,
	string(CategoryScreenwriting):         CategoryScreenwriting,
	string(CategoryStampsAndCoins):        CategoryStampsAndCoins,
	string(CategoryVideoAndComputerGames): CategoryVideoAndComputerGames,
	string(CategoryWoodworking):           CategoryWoodworking,

	string(CategoryHomeAndGarden):             CategoryHomeAndGarden,
	string(CategoryAppliances):                CategoryAppliances,
	string(CategoryEntertaining):              CategoryEntertaining,
	string(CategoryEnvironmentalSafety):       CategoryEnvironmentalSafety,
	string(CategoryGardening):                 CategoryGardening,
	string(CategoryHomeRepair):                CategoryHomeRepair,
	string(CategoryHomeTheater):               CategoryHomeTheater,
	string(CategoryInteriorDecorating):        CategoryInteriorDecorating,
	string(CategoryLandscaping):               CategoryLandscaping,
	string(CategoryRemodelingAndConstruction): CategoryRemodelingAndConstruction,

	string(CategoryLawGovernmentAndPolitics): CategoryLawGovernmentAndPolitics,
	string(CategoryImmigration):              CategoryImmigration,
	string(CategoryLegalIssues):              CategoryLegalIssues,
	string(CategoryUSGovernmentResources):    CategoryUSGovernmentResources,
	string(CategoryPolitics):                 CategoryPolitics,
	string(CategoryCommentary):               CategoryCommentary,
	string(CategoryNews):                     CategoryNews,
	string(CategoryInternationalNews):        CategoryInternationalNews,
	string(CategoryNationalNews):             CategoryNationalNews,
	string(CategoryLocalNews):                CategoryLocalNews,

	string(CategoryPersonalFinance):    CategoryPersonalFinance,
	string(CategoryBeginningInvesting): CategoryBeginningInvesting,
	string(CategoryCreditDebtAndLoans): CategoryCreditDebtAndLoans,
	string(CategoryFinancialNews):      CategoryFinancialNews,
	string(CategoryFinancialPlanning):  CategoryFinancialPlanning,
	string(CategoryHedgeFund):          CategoryHedgeFund,
	string(CategoryInsurance):          CategoryInsurance,
	string(CategoryInvesting):          CategoryInvesting,
	string(CategoryMutualFunds):        CategoryMutualFunds,
	string(CategoryOptions):            CategoryOptions,
	string(CategoryRetirementPlanning): CategoryRetirementPlanning,
	string(CategoryStocks):             CategoryStocks,
	string(CategoryTaxPlanning):        CategoryTaxPlanning,

	string(CategorySociety):        CategorySociety,
	string(CategoryDating):         CategoryDating,
	string(CategoryDivorceSupport): CategoryDivorceSupport,
	string(CategoryGayLife):        CategoryGayLife,
	string(CategoryMarriage):       CategoryMarriage,
	string(CategorySeniorLiving):   CategorySeniorLiving,
	string(CategoryTeens):          CategoryTeens,
	string(CategoryWeddings):       CategoryWeddings,
	string(CategoryEthnicSpecific): CategoryEthnicSpecific,

	string(CategoryScience):             CategoryScience,
	string(CategoryAstrology):           CategoryAstrology,
	string(CategoryBiology):             CategoryBiology,
	string(CategoryChemistry):           CategoryChemistry,
	string(CategoryGeology):             CategoryGeology,
	string(CategoryParanormalPhenomena): CategoryParanormalPhenomena,
	string(CategoryPhysics):             CategoryPhysics,
	string(CategorySpaceAstronomy):      CategorySpaceAstronomy,
	string(CategoryGeography):           CategoryGeography,
	string(CategoryBotany):              CategoryBotany,
	string(CategoryWeather):             CategoryWeather,

	string(CategoryPets):               CategoryPets,
	string(CategoryAquariums):          CategoryAquariums,
	string(CategoryBirds):              CategoryBirds,
	string(CategoryCats):               CategoryCats,
	string(CategoryDogs):               CategoryDogs,
	string(CategoryLargeAnimals):       CategoryLargeAnimals,
	string(CategoryReptiles):           CategoryReptiles,
	string(CategoryVeterinaryMedicine): CategoryVeterinaryMedicine,

	string(CategorySports):              CategorySports,
	string(CategoryAutoRacing):          CategoryAutoRacing,
	string(CategoryBaseball):            CategoryBaseball,
	string(CategoryBicycling):           CategoryBicycling,
	string(CategoryBodybuilding):        CategoryBodybuilding,
	string(CategoryBoxing):              CategoryBoxing,
	string(CategoryCanoeingKayaking):    CategoryCanoeingKayaking,
	string(CategoryCheerleading):        CategoryCheerleading,
	string(CategoryClimbing):            CategoryClimbing,
	string(CategoryCricket):             CategoryCricket,
	string(CategoryFigureSkating):       CategoryFigureSkating,
	string(CategoryFlyFishing):          CategoryFlyFishing,
	string(CategoryFootball):            CategoryFootball,
	string(CategoryFreshwaterFishing):   CategoryFreshwaterFishing,
	string(CategoryGameAndFish):         CategoryGameAndFish,
	string(CategoryGolf):                CategoryGolf,
	string(CategoryHorseRacing):         CategoryHorseRacing,
	string(CategoryHorses):              CategoryHorses,
	string(CategoryHuntingShooting):     CategoryHuntingShooting,
	string(CategoryInlineSkating):       CategoryInlineSkating,
	string(CategoryMartialArts):         CategoryMartialArts,
	string(CategoryMountainBiking):      CategoryMountainBiking,
	string(CategoryNASCARRacing):        CategoryNASCARRacing,
	string(CategoryOlympics):            CategoryOlympics,
	string(CategoryPaintball):           CategoryPaintball,
	string(CategoryPowerAndMotorcycles): CategoryPowerAndMotorcycles,
	string(CategoryProBasketball):       CategoryProBasketball,
	string(CategoryProIceHockey):        CategoryProIceHockey,
	string(CategoryRodeo):               CategoryRodeo,
	string(CategoryRugby):               CategoryRugby,
	string(CategoryRunningJogging):      CategoryRunningJogging,
	string(CategorySailing):             CategorySailing,
	string(CategorySaltwaterFishing):    CategorySaltwaterFishing,
	string(CategoryScubaDiving):         CategoryScubaDiving,
	string(CategorySkateboarding):       CategorySkateboarding,
	string(CategorySkiing):              CategorySkiing,
	string(CategorySnowboarding):        CategorySnowboarding,
	string(CategorySurfingBodyboarding): CategorySurfingBodyboarding,
	string(CategorySwimming):            CategorySwimming,
	string(CategoryTableTennisPingPong): CategoryTableTennisPingPong,
	string(CategoryTennis):              CategoryTennis,
	string(CategoryVolleyball):          CategoryVolleyball,
	string(CategoryWalking):             CategoryWalking,
	string(CategoryWaterskiWakeboard):   CategoryWaterskiWakeboard,
	string(CategoryWorldSoccer):         CategoryWorldSoccer,

	string(CategoryStyleAndFashion): CategoryStyleAndFashion,
	string(CategoryBeauty):          CategoryBeauty,
	string(CategoryBodyArt):         CategoryBodyArt,
	string(CategoryFashion):         CategoryFashion,
	string(CategoryJewelry):         CategoryJewelry,
	string(CategoryClothing):        CategoryClothing,
	string(CategoryAccessories):     CategoryAccessories,

	string(CategoryTechnologyAndComputing): CategoryTechnologyAndComputing,
	string(Category3DGraphics):             Category3DGraphics,
	string(CategoryAnimation):              CategoryAnimation,
	string(CategoryAntivirusSoftware):      CategoryAntivirusSoftware,
	string(CategoryCCPlusPlus):             CategoryCCPlusPlus,
	string(CategoryCamerasAndCamcorders):   CategoryCamerasAndCamcorders,
	string(CategoryCellPhones):             CategoryCellPhones,
	string(CategoryComputerCertification):  CategoryComputerCertification,
	string(CategoryComputerNetworking):     CategoryComputerNetworking,
	string(CategoryComputerPeripherals):    CategoryComputerPeripherals,
	string(CategoryComputerReviews):        CategoryComputerReviews,
	string(CategoryDataCenters):            CategoryDataCenters,
	string(CategoryDatabases):              CategoryDatabases,
	string(CategoryDesktopPublishing):      CategoryDesktopPublishing,
	string(CategoryDesktopVideo):           CategoryDesktopVideo,
	string(CategoryEmail):                  CategoryEmail,
	string(CategoryGraphicsSoftware):       CategoryGraphicsSoftware,
	string(CategoryHomeVideoDVD):           CategoryHomeVideoDVD,
	string(CategoryInternetTechnology):     CategoryInternetTechnology,
	string(CategoryJava):                   CategoryJava,
	string(CategoryJavaScript):             CategoryJavaScript,
	string(CategoryMacSupport):             CategoryMacSupport,
	string(CategoryMP3MIDI):                CategoryMP3MIDI,
	string(CategoryNetConferencing):        CategoryNetConferencing,
	string(CategoryNetforBeginners):        CategoryNetforBeginners,
	string(CategoryNetworkSecurity):        CategoryNetworkSecurity,
	string(CategoryPalmtopsPDAs):           CategoryPalmtopsPDAs,
	string(CategoryPCSupport):              CategoryPCSupport,
	string(CategoryPortable):               CategoryPortable,
	string(CategoryEntertainment):          CategoryEntertainment,
	string(CategorySharewareFreeware):      CategorySharewareFreeware,
	string(CategoryUnix):                   CategoryUnix,
	string(CategoryVisualBasic):            CategoryVisualBasic,
	string(CategoryWebClipArt):             CategoryWebClipArt,
	string(CategoryWebDesignHTML):          CategoryWebDesignHTML,
	string(CategoryWebSearch):              CategoryWebSearch,
	string(CategoryWindows):                CategoryWindows,

	string(CategoryTravel):                  CategoryTravel,
	string(CategoryAdventureTravel):         CategoryAdventureTravel,
	string(CategoryAfrica):                  CategoryAfrica,
	string(CategoryAirTravel):               CategoryAirTravel,
	string(CategoryAustraliaAndNewZealand):  CategoryAustraliaAndNewZealand,
	string(CategoryBedAndBreakfasts):        CategoryBedAndBreakfasts,
	string(CategoryBudgetTravel):            CategoryBudgetTravel,
	string(CategoryBusinessTravel):          CategoryBusinessTravel,
	string(CategoryByUSLocale):              CategoryByUSLocale,
	string(CategoryCamping):                 CategoryCamping,
	string(CategoryCanada):                  CategoryCanada,
	string(CategoryCaribbean):               CategoryCaribbean,
	string(CategoryCruises):                 CategoryCruises,
	string(CategoryEasternEurope):           CategoryEasternEurope,
	string(CategoryEurope):                  CategoryEurope,
	string(CategoryFrance):                  CategoryFrance,
	string(CategoryGreece):                  CategoryGreece,
	string(CategoryHoneymoonsGetaways):      CategoryHoneymoonsGetaways,
	string(CategoryHotels):                  CategoryHotels,
	string(CategoryItaly):                   CategoryItaly,
	string(CategoryJapan):                   CategoryJapan,
	string(CategoryMexicoAndCentralAmerica): CategoryMexicoAndCentralAmerica,
	string(CategoryNationalParks):           CategoryNationalParks,
	string(CategorySouthAmerica):            CategorySouthAmerica,
	string(CategorySpas):                    CategorySpas,
	string(CategoryThemeParks):              CategoryThemeParks,
	string(CategoryTravelingwithKids):       CategoryTravelingwithKids,
	string(CategoryUnitedKingdom):           CategoryUnitedKingdom,

	string(CategoryRealEstate):         CategoryRealEstate,
	string(CategoryApartments):         CategoryApartments,
	string(CategoryArchitects):         CategoryArchitects,
	string(CategoryBuyingSellingHomes): CategoryBuyingSellingHomes,

	string(CategoryShopping):            CategoryShopping,
	string(CategoryContestsAndFreebies): CategoryContestsAndFreebies,
	string(CategoryCouponing):           CategoryCouponing,
	string(CategoryComparison):          CategoryComparison,
	string(CategoryEngines):             CategoryEngines,

	string(CategoryReligionAndSpirituality): CategoryReligionAndSpirituality,
	string(CategoryAlternativeReligions):    CategoryAlternativeReligions,
	string(CategoryAtheismAgnosticism):      CategoryAtheismAgnosticism,
	string(CategoryBuddhism):                CategoryBuddhism,
	string(CategoryCatholicism):             CategoryCatholicism,
	string(CategoryChristianity):            CategoryChristianity,
	string(CategoryHinduism):                CategoryHinduism,
	string(CategoryIslam):                   CategoryIslam,
	string(CategoryJudaism):                 CategoryJudaism,
	string(CategoryLatterDaySaints):         CategoryLatterDaySaints,
	string(CategoryPaganWiccan):             CategoryPaganWiccan,

	string(CategoryUncategorized): CategoryUncategorized,

	string(CategoryNonStandardContent):             CategoryNonStandardContent,
	string(CategoryUnmoderatedUGC):                 CategoryUnmoderatedUGC,
	string(CategoryExtremeGraphicExplicitViolence): CategoryExtremeGraphicExplicitViolence,
	string(CategoryPornography):                    CategoryPornography,
	string(CategoryProfaneContent):                 CategoryProfaneContent,
	string(CategoryHateContent):                    CategoryHateContent,
	string(CategoryUnderConstruction):              CategoryUnderConstruction,
	string(CategoryIncentivized):                   CategoryIncentivized,

	string(CategoryIllegalContent):        CategoryIllegalContent,
	string(CategoryIllegalContent2):       CategoryIllegalContent2,
	string(CategoryWarez):                 CategoryWarez,
	string(CategorySpywareMalware):        CategorySpywareMalware,
	string(CategoryCopyrightInfringement): CategoryCopyrightInfringement,
}

// MarshalCategory converts a string into an Category, and returns an error if
// there is no corresponding Category.
func MarshalCategory(value string) (Category, error) {
	cat, ok := categoryNameToValue[value]
	if !ok {
		return "", fmt.Errorf(`can't marshal "%s" to Category`, value)
	}
	return cat, nil
}
