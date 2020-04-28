package openrtb

// VolumeNormalizationMode struct represents the types of volume normalization modes, typically for audio.
// See OpenRTB Sec 5.17.
type VolumeNormalizationMode int

// Possible values from the OpenRTB spec.
const (
	VolumeNormalizationModeNone     VolumeNormalizationMode = 0 // None
	VolumeNormalizationModeAverage  VolumeNormalizationMode = 1 // Ad Volume Average Normalized to Content
	VolumeNormalizationModePeak     VolumeNormalizationMode = 2 // Ad Volume Peak Normalized to Content
	VolumeNormalizationModeLoudness VolumeNormalizationMode = 3 // Ad Loudness Normalized to Content
	VolumeNormalizationModeCustom   VolumeNormalizationMode = 4 // Custom Volume Normalization
)
