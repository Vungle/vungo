package openrtb

// VolumeNormalizationMode struct represents the types of volume normalization modes, typically for audio.
// See OpenRTB Sec 5.17.
type VolumeNormalizationMode int

// Possible values from the OpenRTB spec.
const (
	VolumeNormalizationModeNone     VolumeNormalizationMode = 0
	VolumeNormalizationModeAverage  VolumeNormalizationMode = 1
	VolumeNormalizationModePeak     VolumeNormalizationMode = 2
	VolumeNormalizationModeLoudness VolumeNormalizationMode = 3
	VolumeNormalizationModeCustom   VolumeNormalizationMode = 4
)
