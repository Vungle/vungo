package openrtb

// This file contains constants defined in the Advertising Common Object Model.

// AgentType identifies the user agent types a user identifier is from.
type AgentType int64

// Agent types describing where the user agent is from.
//
// Values of 500+ hold vendor-specific codes.
const (
	AgentTypeWeb    AgentType = 1 // An ID which is tied to a specific web browser or device (cookie-based, probabilistic, or other).
	AgentTypeApp    AgentType = 2 // In-app impressions, which will typically contain a type of device ID (or rather, the privacy-compliant versions of device IDs).
	AgentTypePerson AgentType = 3 // A person-based ID, i.e., that is the same across devices.
)

// MatchMethod represents various ways an ID could be matched to an ad request, and if they pertain to a single property or app.
// It should be used on conjunction with the mm attribute in Object: EID of OpenRTB 2.x.
type MatchMethod int64

// Various ways an ID could be matched to an ad request
const (
	MatchMethodUnknown           MatchMethod = 0
	MatchMethodNoMatch           MatchMethod = 1 // No matching has occurred. The associated ID came directly from a 3rd-party cookie or OS-provided resettable device ID for advertising (IFA).
	MatchMethodBrowserCookieSync MatchMethod = 2 // Real time cookie sync as described in Appendix C (Cookie Based ID Syncing) of OpenRTB 2.x.
	MatchMethodAuthenticated     MatchMethod = 3 // ID match was based on user authentication such as an email login or hashed PII.
	MatchMethodObserved          MatchMethod = 4 // ID match was based on a 1st party observation, but without user authentication (e.g. GUID, SharedID, Session IDs, CHIPS or other 1st party cookies contained in localStorage).
	MatchMethodInference         MatchMethod = 5 // ID match was inferred from linkage based on non-authenticated features across multiple browsers or devices (e.g. IP address and/or UserAgent).
)
