package openrtb

// StartDelay represents various options for the video or audio start delay (5.12). If the start delay
// value is greater than 0, then the position is mid-roll and the value indicates the start delay.
// See OpenRTB 2.5 Sec 5.12.
type StartDelay int

// StartDelay enum values
const (
	// > 0 Mid-Roll (value indicates start delay in second)

	StartDelayPreRoll         StartDelay = 0  // StartDelayPreRoll represents Pre-Roll.
	StartDelayGenericMidRoll  StartDelay = -1 // StartDelayGenericMidRoll represents Generic Mid-Roll.
	StartDelayGenericPostRoll StartDelay = -2 // StartDelayGenericPostRoll represents Generic Post-Roll.
)

// Copy do deep copy of StartDelay.
func (nb *StartDelay) Copy() *StartDelay {
	if nb == nil {
		return nil
	}
	nbCopy := *nb
	return &nbCopy
}
