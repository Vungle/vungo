package openrtb

// VolumeNormalizationMode struct represents the types of volume normalization modes, typically for audio.
// See OpenRTB Sec 5.17.
type VolumeNormalizationMode int

// Possible values from the OpenRTB spec.
const (
	VolumeNormalizationModeNone VolumeNormalizationMode = iota + 1
	VolumeNormalizationModeAverage
	VolumeNormalizationModePeak
	VolumeNormalizationModeLoudness
	VolumeNormalizationModeCustom
)
