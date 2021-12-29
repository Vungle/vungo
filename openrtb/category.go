package openrtb

// Possible values of IAB content categories.
const (
	CategoryArtsAndEntertainment string = "IAB1"
	CategoryBooksAndLiterature   string = "IAB1-1"
	CategoryCelebrityFanGossip   string = "IAB1-2"
	CategoryFineArt              string = "IAB1-3"
	CategoryHumor                string = "IAB1-4"
	CategoryMovies               string = "IAB1-5"
	CategoryMusic                string = "IAB1-6"
	CategoryTelevision           string = "IAB1-7"

	CategoryAutomotive           string = "IAB2"
	CategoryAutoParts            string = "IAB2-1"
	CategoryAutoRepair           string = "IAB2-2"
	CategoryBuyingSellingCars    string = "IAB2-3"
	CategoryCarCulture           string = "IAB2-4"
	CategoryCertifiedPreOwned    string = "IAB2-5"
	CategoryConvertible          string = "IAB2-6"
	CategoryCoupe                string = "IAB2-7"
	CategoryCrossover            string = "IAB2-8"
	CategoryDiesel               string = "IAB2-9"
	CategoryElectricVehicle      string = "IAB2-10"
	CategoryHatchback            string = "IAB2-11"
	CategoryHybrid               string = "IAB2-12"
	CategoryLuxury               string = "IAB2-13"
	CategoryMinivan              string = "IAB2-14"
	CategoryMororcycles          string = "IAB2-15"
	CategoryOffRoadVehicles      string = "IAB2-16"
	CategoryPerformanceVehicles  string = "IAB2-17"
	CategoryPickup               string = "IAB2-18"
	CategoryRoadSideAssistance   string = "IAB2-19"
	CategorySedan                string = "IAB2-20"
	CategoryTrucksAndAccessories string = "IAB2-21"
	CategoryVintageCars          string = "IAB2-22"
	CategoryWagon                string = "IAB2-23"

	CategoryBusiness          string = "IAB3"
	CategoryAdvertising       string = "IAB3-1"
	CategoryAgriculture       string = "IAB3-2"
	CategoryBiotechBiomedical string = "IAB3-3"
	CategoryBusinessSoftware  string = "IAB3-4"
	CategoryConstruction      string = "IAB3-5"
	CategoryForestry          string = "IAB3-6"
	CategoryGovernment        string = "IAB3-7"
	CategoryGreenSolutions    string = "IAB3-8"
	CategoryHumanResources    string = "IAB3-9"
	CategoryLogistics         string = "IAB3-10"
	CategoryMarketing         string = "IAB3-11"
	CategoryMetals            string = "IAB3-12"

	CategoryCareers             string = "IAB4"
	CategoryCareerPlanning      string = "IAB4-1"
	CategoryCollege             string = "IAB4-2"
	CategoryFinancialAid        string = "IAB4-3"
	CategoryJobFairs            string = "IAB4-4"
	CategoryJobSearch           string = "IAB4-5"
	CategoryResumeWritingAdvice string = "IAB4-6"
	CategoryNursing             string = "IAB4-7"
	CategoryScholarships        string = "IAB4-8"
	CategoryTelecommuting       string = "IAB4-9"
	CategoryUSMilitary          string = "IAB4-10"
	CategoryCareerAdvice        string = "IAB4-11"

	CategoryEducation              string = "IAB5"
	Category712Education           string = "IAB5-1"
	CategoryAdultEducation         string = "IAB5-2"
	CategoryArtHistory             string = "IAB5-3"
	CategoryColledgeAdministration string = "IAB5-4"
	CategoryCollegeLife            string = "IAB5-5"
	CategoryDistanceLearning       string = "IAB5-6"
	CategoryEnglishAsA2ndLanguage  string = "IAB5-7"
	CategoryLanguageLearning       string = "IAB5-8"
	CategoryGraduateSchool         string = "IAB5-9"
	CategoryHomeschooling          string = "IAB5-10"
	CategoryHomeworkStudyTips      string = "IAB5-11"
	CategoryK6Educators            string = "IAB5-12"
	CategoryPrivateSchool          string = "IAB5-13"
	CategorySpecialEducation       string = "IAB5-14"
	CategoryStudyingBusiness       string = "IAB5-15"

	CategoryFamilyAndParenting string = "IAB6"
	CategoryAdoption           string = "IAB6-1"
	CategoryBabiesAndToddlers  string = "IAB6-2"
	CategoryDaycarePreSchool   string = "IAB6-3"
	CategoryFamilyInternet     string = "IAB6-4"
	CategoryParentingK6Kids    string = "IAB6-5"
	CategoryParentingTeens     string = "IAB6-6"
	CategoryPregnancy          string = "IAB6-7"
	CategorySpecialNeedsKids   string = "IAB6-8"
	CategoryEldercare          string = "IAB6-9"

	CategoryHealthAndFitness       string = "IAB7"
	CategoryExercise               string = "IAB7-1"
	CategoryADD                    string = "IAB7-2"
	CategoryAIDSHIV                string = "IAB7-3"
	CategoryAllergies              string = "IAB7-4"
	CategoryAlternativeMedicine    string = "IAB7-5"
	CategoryArthritis              string = "IAB7-6"
	CategoryAsthma                 string = "IAB7-7"
	CategoryAutismPDD              string = "IAB7-8"
	CategoryBipolarDisorder        string = "IAB7-9"
	CategoryBrainTumor             string = "IAB7-10"
	CategoryCancer                 string = "IAB7-11"
	CategoryCholesterol            string = "IAB7-12"
	CategoryChronicFatigueSyndrome string = "IAB7-13"
	CategoryChronicPain            string = "IAB7-14"
	CategoryColdAndFlu             string = "IAB7-15"
	CategoryDeafness               string = "IAB7-16"
	CategoryDentalCare             string = "IAB7-17"
	CategoryDepression             string = "IAB7-18"
	CategoryDermatology            string = "IAB7-19"
	CategoryDiabetes               string = "IAB7-20"
	CategoryEpilepsy               string = "IAB7-21"
	CategoryGERDAcidReflux         string = "IAB7-22"
	CategoryHeadachesMigraines     string = "IAB7-23"
	CategoryHeartDisease           string = "IAB7-24"
	CategoryHerbsforHealth         string = "IAB7-25"
	CategoryHolisticHealing        string = "IAB7-26"
	CategoryIBSCrohnsDisease       string = "IAB7-27"
	CategoryIncestAbuseSupport     string = "IAB7-28"
	CategoryIncontinence           string = "IAB7-29"
	CategoryInfertility            string = "IAB7-30"
	CategoryMensHealth             string = "IAB7-31"
	CategoryNutrition              string = "IAB7-32"
	CategoryOrthopedics            string = "IAB7-33"
	CategoryPanicAnxietyDisorders  string = "IAB7-34"
	CategoryPediatrics             string = "IAB7-35"
	CategoryPhysicalTherapy        string = "IAB7-36"
	CategoryPsychologyPsychiatry   string = "IAB7-37"
	CategorySeniorHealth           string = "IAB7-38"
	CategorySexuality              string = "IAB7-39"
	CategorySleepDisorders         string = "IAB7-40"
	CategorySmokingCessation       string = "IAB7-41"
	CategorySubstanceAbuse         string = "IAB7-42"
	CategoryThyroidDisease         string = "IAB7-43"
	CategoryWeightLoss             string = "IAB7-44"
	CategoryWomensHealth           string = "IAB7-45"

	CategoryFoodAndDrink         string = "IAB8"
	CategoryAmericanCuisine      string = "IAB8-1"
	CategoryBarbecuesAndGrilling string = "IAB8-2"
	CategoryCajunCreole          string = "IAB8-3"
	CategoryChineseCuisine       string = "IAB8-4"
	CategoryCocktailsBeer        string = "IAB8-5"
	CategoryCoffeeTea            string = "IAB8-6"
	CategoryCuisineSpecific      string = "IAB8-7"
	CategoryDessertsAndBaking    string = "IAB8-8"
	CategoryDiningOut            string = "IAB8-9"
	CategoryFoodAllergies        string = "IAB8-10"
	CategoryFrenchCuisine        string = "IAB8-11"
	CategoryHealthLowFatCooking  string = "IAB8-12"
	CategoryItalianCuisine       string = "IAB8-13"
	CategoryJapaneseCuisine      string = "IAB8-14"
	CategoryMexicanCuisine       string = "IAB8-15"
	CategoryVegan                string = "IAB8-16"
	CategoryVegetarian           string = "IAB8-17"
	CategoryWine                 string = "IAB8-18"

	CategoryHobbiesAndInterests   string = "IAB9"
	CategoryArtTechnology         string = "IAB9-1"
	CategoryArtsAndCrafts         string = "IAB9-2"
	CategoryBeadwork              string = "IAB9-3"
	CategoryBirdwatching          string = "IAB9-4"
	CategoryBoardGamesPuzzles     string = "IAB9-5"
	CategoryCandleAndSoapMaking   string = "IAB9-6"
	CategoryCardGames             string = "IAB9-7"
	CategoryChess                 string = "IAB9-8"
	CategoryCigars                string = "IAB9-9"
	CategoryCollecting            string = "IAB9-10"
	CategoryComicBooks            string = "IAB9-11"
	CategoryDrawingSketching      string = "IAB9-12"
	CategoryFreelanceWriting      string = "IAB9-13"
	CategoryGenealogy             string = "IAB9-14"
	CategoryGettingPublished      string = "IAB9-15"
	CategoryGuitar                string = "IAB9-16"
	CategoryHomeRecording         string = "IAB9-17"
	CategoryInvestorsAndPatents   string = "IAB9-18"
	CategoryJewelryMaking         string = "IAB9-19"
	CategoryMagicAndIllusion      string = "IAB9-20"
	CategoryNeedlework            string = "IAB9-21"
	CategoryPainting              string = "IAB9-22"
	CategoryPhotography           string = "IAB9-23"
	CategoryRadio                 string = "IAB9-24"
	CategoryRoleplayingGames      string = "IAB9-25"
	CategorySciFiAndFantasy       string = "IAB9-26"
	CategoryScrapbooking          string = "IAB9-27"
	CategoryScreenwriting         string = "IAB9-28"
	CategoryStampsAndCoins        string = "IAB9-29"
	CategoryVideoAndComputerGames string = "IAB9-30"
	CategoryWoodworking           string = "IAB9-31"

	CategoryHomeAndGarden             string = "IAB10"
	CategoryAppliances                string = "IAB10-1"
	CategoryEntertaining              string = "IAB10-2"
	CategoryEnvironmentalSafety       string = "IAB10-3"
	CategoryGardening                 string = "IAB10-4"
	CategoryHomeRepair                string = "IAB10-5"
	CategoryHomeTheater               string = "IAB10-6"
	CategoryInteriorDecorating        string = "IAB10-7"
	CategoryLandscaping               string = "IAB10-8"
	CategoryRemodelingAndConstruction string = "IAB10-9"

	CategoryLawGovernmentAndPolitics string = "IAB11"
	CategoryImmigration              string = "IAB11-1"
	CategoryLegalIssues              string = "IAB11-2"
	CategoryUSGovernmentResources    string = "IAB11-3"
	CategoryPolitics                 string = "IAB11-4"
	CategoryCommentary               string = "IAB11-5"
	CategoryNews                     string = "IAB12"
	CategoryInternationalNews        string = "IAB12-1"
	CategoryNationalNews             string = "IAB12-2"
	CategoryLocalNews                string = "IAB12-3"

	CategoryPersonalFinance    string = "IAB13"
	CategoryBeginningInvesting string = "IAB13-1"
	CategoryCreditDebtAndLoans string = "IAB13-2"
	CategoryFinancialNews      string = "IAB13-3"
	CategoryFinancialPlanning  string = "IAB13-4"
	CategoryHedgeFund          string = "IAB13-5"
	CategoryInsurance          string = "IAB13-6"
	CategoryInvesting          string = "IAB13-7"
	CategoryMutualFunds        string = "IAB13-8"
	CategoryOptions            string = "IAB13-9"
	CategoryRetirementPlanning string = "IAB13-10"
	CategoryStocks             string = "IAB13-11"
	CategoryTaxPlanning        string = "IAB13-12"

	CategorySociety        string = "IAB14"
	CategoryDating         string = "IAB14-1"
	CategoryDivorceSupport string = "IAB14-2"
	CategoryGayLife        string = "IAB14-3"
	CategoryMarriage       string = "IAB14-4"
	CategorySeniorLiving   string = "IAB14-5"
	CategoryTeens          string = "IAB14-6"
	CategoryWeddings       string = "IAB14-7"
	CategoryEthnicSpecific string = "IAB14-8"

	CategoryScience             string = "IAB15"
	CategoryAstrology           string = "IAB15-1"
	CategoryBiology             string = "IAB15-2"
	CategoryChemistry           string = "IAB15-3"
	CategoryGeology             string = "IAB15-4"
	CategoryParanormalPhenomena string = "IAB15-5"
	CategoryPhysics             string = "IAB15-6"
	CategorySpaceAstronomy      string = "IAB15-7"
	CategoryGeography           string = "IAB15-8"
	CategoryBotany              string = "IAB15-9"
	CategoryWeather             string = "IAB15-10"

	CategoryPets               string = "IAB16"
	CategoryAquariums          string = "IAB16-1"
	CategoryBirds              string = "IAB16-2"
	CategoryCats               string = "IAB16-3"
	CategoryDogs               string = "IAB16-4"
	CategoryLargeAnimals       string = "IAB16-5"
	CategoryReptiles           string = "IAB16-6"
	CategoryVeterinaryMedicine string = "IAB16-7"

	CategorySports              string = "IAB17"
	CategoryAutoRacing          string = "IAB17-1"
	CategoryBaseball            string = "IAB17-2"
	CategoryBicycling           string = "IAB17-3"
	CategoryBodybuilding        string = "IAB17-4"
	CategoryBoxing              string = "IAB17-5"
	CategoryCanoeingKayaking    string = "IAB17-6"
	CategoryCheerleading        string = "IAB17-7"
	CategoryClimbing            string = "IAB17-8"
	CategoryCricket             string = "IAB17-9"
	CategoryFigureSkating       string = "IAB17-10"
	CategoryFlyFishing          string = "IAB17-11"
	CategoryFootball            string = "IAB17-12"
	CategoryFreshwaterFishing   string = "IAB17-13"
	CategoryGameAndFish         string = "IAB17-14"
	CategoryGolf                string = "IAB17-15"
	CategoryHorseRacing         string = "IAB17-16"
	CategoryHorses              string = "IAB17-17"
	CategoryHuntingShooting     string = "IAB17-18"
	CategoryInlineSkating       string = "IAB17-19"
	CategoryMartialArts         string = "IAB17-20"
	CategoryMountainBiking      string = "IAB17-21"
	CategoryNASCARRacing        string = "IAB17-22"
	CategoryOlympics            string = "IAB17-23"
	CategoryPaintball           string = "IAB17-24"
	CategoryPowerAndMotorcycles string = "IAB17-25"
	CategoryProBasketball       string = "IAB17-26"
	CategoryProIceHockey        string = "IAB17-27"
	CategoryRodeo               string = "IAB17-28"
	CategoryRugby               string = "IAB17-29"
	CategoryRunningJogging      string = "IAB17-30"
	CategorySailing             string = "IAB17-31"
	CategorySaltwaterFishing    string = "IAB17-32"
	CategoryScubaDiving         string = "IAB17-33"
	CategorySkateboarding       string = "IAB17-34"
	CategorySkiing              string = "IAB17-35"
	CategorySnowboarding        string = "IAB17-36"
	CategorySurfingBodyboarding string = "IAB17-37"
	CategorySwimming            string = "IAB17-38"
	CategoryTableTennisPingPong string = "IAB17-39"
	CategoryTennis              string = "IAB17-40"
	CategoryVolleyball          string = "IAB17-41"
	CategoryWalking             string = "IAB17-42"
	CategoryWaterskiWakeboard   string = "IAB17-43"
	CategoryWorldSoccer         string = "IAB17-44"

	CategoryStyleAndFashion string = "IAB18"
	CategoryBeauty          string = "IAB18-1"
	CategoryBodyArt         string = "IAB18-2"
	CategoryFashion         string = "IAB18-3"
	CategoryJewelry         string = "IAB18-4"
	CategoryClothing        string = "IAB18-5"
	CategoryAccessories     string = "IAB18-6"

	CategoryTechnologyAndComputing string = "IAB19"
	Category3DGraphics             string = "IAB19-1"
	CategoryAnimation              string = "IAB19-2"
	CategoryAntivirusSoftware      string = "IAB19-3"
	CategoryCCPlusPlus             string = "IAB19-4"
	CategoryCamerasAndCamcorders   string = "IAB19-5"
	CategoryCellPhones             string = "IAB19-6"
	CategoryComputerCertification  string = "IAB19-7"
	CategoryComputerNetworking     string = "IAB19-8"
	CategoryComputerPeripherals    string = "IAB19-9"
	CategoryComputerReviews        string = "IAB19-10"
	CategoryDataCenters            string = "IAB19-11"
	CategoryDatabases              string = "IAB19-12"
	CategoryDesktopPublishing      string = "IAB19-13"
	CategoryDesktopVideo           string = "IAB19-14"
	CategoryEmail                  string = "IAB19-15"
	CategoryGraphicsSoftware       string = "IAB19-16"
	CategoryHomeVideoDVD           string = "IAB19-17"
	CategoryInternetTechnology     string = "IAB19-18"
	CategoryJava                   string = "IAB19-19"
	CategoryJavaScript             string = "IAB19-20"
	CategoryMacSupport             string = "IAB19-21"
	CategoryMP3MIDI                string = "IAB19-22"
	CategoryNetConferencing        string = "IAB19-23"
	CategoryNetforBeginners        string = "IAB19-24"
	CategoryNetworkSecurity        string = "IAB19-25"
	CategoryPalmtopsPDAs           string = "IAB19-26"
	CategoryPCSupport              string = "IAB19-27"
	CategoryPortable               string = "IAB19-28"
	CategoryEntertainment          string = "IAB19-29"
	CategorySharewareFreeware      string = "IAB19-30"
	CategoryUnix                   string = "IAB19-31"
	CategoryVisualBasic            string = "IAB19-32"
	CategoryWebClipArt             string = "IAB19-33"
	CategoryWebDesignHTML          string = "IAB19-34"
	CategoryWebSearch              string = "IAB19-35"
	CategoryWindows                string = "IAB19-36"

	CategoryTravel                  string = "IAB20"
	CategoryAdventureTravel         string = "IAB20-1"
	CategoryAfrica                  string = "IAB20-2"
	CategoryAirTravel               string = "IAB20-3"
	CategoryAustraliaAndNewZealand  string = "IAB20-4"
	CategoryBedAndBreakfasts        string = "IAB20-5"
	CategoryBudgetTravel            string = "IAB20-6"
	CategoryBusinessTravel          string = "IAB20-7"
	CategoryByUSLocale              string = "IAB20-8"
	CategoryCamping                 string = "IAB20-9"
	CategoryCanada                  string = "IAB20-10"
	CategoryCaribbean               string = "IAB20-11"
	CategoryCruises                 string = "IAB20-12"
	CategoryEasternEurope           string = "IAB20-13"
	CategoryEurope                  string = "IAB20-14"
	CategoryFrance                  string = "IAB20-15"
	CategoryGreece                  string = "IAB20-16"
	CategoryHoneymoonsGetaways      string = "IAB20-17"
	CategoryHotels                  string = "IAB20-18"
	CategoryItaly                   string = "IAB20-19"
	CategoryJapan                   string = "IAB20-20"
	CategoryMexicoAndCentralAmerica string = "IAB20-21"
	CategoryNationalParks           string = "IAB20-22"
	CategorySouthAmerica            string = "IAB20-23"
	CategorySpas                    string = "IAB20-24"
	CategoryThemeParks              string = "IAB20-25"
	CategoryTravelingwithKids       string = "IAB20-26"
	CategoryUnitedKingdom           string = "IAB20-27"

	CategoryRealEstate         string = "IAB21"
	CategoryApartments         string = "IAB21-1"
	CategoryArchitects         string = "IAB21-2"
	CategoryBuyingSellingHomes string = "IAB21-3"

	CategoryShopping            string = "IAB22"
	CategoryContestsAndFreebies string = "IAB22-1"
	CategoryCouponing           string = "IAB22-2"
	CategoryComparison          string = "IAB22-3"
	CategoryEngines             string = "IAB22-4"

	CategoryReligionAndSpirituality string = "IAB23"
	CategoryAlternativeReligions    string = "IAB23-1"
	CategoryAtheismAgnosticism      string = "IAB23-2"
	CategoryBuddhism                string = "IAB23-3"
	CategoryCatholicism             string = "IAB23-4"
	CategoryChristianity            string = "IAB23-5"
	CategoryHinduism                string = "IAB23-6"
	CategoryIslam                   string = "IAB23-7"
	CategoryJudaism                 string = "IAB23-8"
	CategoryLatterDaySaints         string = "IAB23-9"
	CategoryPaganWiccan             string = "IAB23-10"

	CategoryUncategorized string = "IAB24"

	CategoryNonStandardContent             string = "IAB25"
	CategoryUnmoderatedUGC                 string = "IAB25-1"
	CategoryExtremeGraphicExplicitViolence string = "IAB25-2"
	CategoryPornography                    string = "IAB25-3"
	CategoryProfaneContent                 string = "IAB25-4"
	CategoryHateContent                    string = "IAB25-5"
	CategoryUnderConstruction              string = "IAB25-6"
	CategoryIncentivized                   string = "IAB25-7"

	CategoryIllegalContent        string = "IAB26"
	CategoryIllegalContent2       string = "IAB26-1"
	CategoryWarez                 string = "IAB26-2"
	CategorySpywareMalware        string = "IAB26-3"
	CategoryCopyrightInfringement string = "IAB26-4"
)
