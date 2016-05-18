package openrtb

type AdPosition int

const (
	AD_POSITION_UNKNOWN AdPosition = iota
	AD_POSITION_ABOVE_FOLD
	_ // Deprecated: Screen position is not used in OpenRTB 2.3.1.
	AD_POSITION_BELOW_FOLD
	AD_POSITION_HEADER
	AD_POSITION_FOOTER
	AD_POSITION_SIDEBAR
	AD_POSITION_FULLSCREEN
)
