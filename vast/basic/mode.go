package vastbasic

// Mode type represents the various ways of how a companion ad can be displayed.
type Mode string

// Enumeration of possible modes of how a companion ad can be displayed.
const (
	ModeAll  Mode = "all"  // All companion ads must be displayed.
	ModeAny  Mode = "any"  // At least one companion ad must be displayed.
	ModeNone Mode = "none" // Companion ad display is optional.
)
