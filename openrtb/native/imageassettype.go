package native

// ImageAssetType indicates the image asset type.
// Below is a list of common image asset element types of native advertising at the time of writing this spec.
// This list is non-exhaustive and intended to be extended by the buyers and sellers as the format evolves.
// An implementing exchange may not support all asset variants or may introduce new ones unique to that system.
// In order to facilitate adoption, recommendations are made for both minimum sizes and aspect ratios.
// We speak here of 'minimum maximum height' or ‘max height of at least’, which means the SSP should support a max
// height of at least this value. They are free to support larger, but the DSP knows that if they have an image of this
// size it will be accepted. Note that SSPs will be responsible for sizing image to exact size if min-max-height
// framework is used; exact size may not be available at bid request time. Width is calculated from the 3 supported
// aspect ratios. Note we merged the prior overlapping type 1 and type 2 as just type 1 - to be used for app icon, brand
// logo, or similar.
//
// See OpenRTB Native 1.2 Sec 7.5 Image Asset Types
type ImageAssetType int64

const (
	// ImageAssetTypeIcon is the type for icon.
	// Type ID:
	//   1
	// Name:
	//   Icon
	// Description:
	//   Icon image
	// Recommendations:
	//   Optional. max height: at least 50. aspect ratio: 1:1
	ImageAssetTypeIcon ImageAssetType = 1

	// ImageAssetTypeMain is the type for main image.
	// Type ID:
	//   3
	// Name:
	//   Main
	// Description:
	//   Large image preview for the ad.
	// Recommendations:
	//   At least one of 2 size variants required:
	//	 Small Variant:
	//	 	max height: at least 200,
	//	 	max width: at least 200, 267, or 382,
	//	 	aspect ratio: 1:1, 4:3, or 1.91:1
	//   Large Variant:
	//      max height: at least 627
	//      max width: at least 627, 836, or 1198
	//      aspect ratio: 1:1, 4:3, or 1.91:1
	ImageAssetTypeMain ImageAssetType = 3

	// Type ID:
	//   500+
	// Name:
	//   XXX
	// Description:
	//   Reserved for Exchange specific usage numbered above 500.
	// Recommendations:
	//	 No recommendations
)
