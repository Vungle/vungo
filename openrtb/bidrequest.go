package openrtb

import (
	"fmt"

	"github.com/Vungle/vungo/internal/util"
)

// BidRequest is the top-level bid request object. It contains a globally unique
// bid request or auction ID.
// This id attribute is required as is at least one impression object (Section
// 3.2.4).
// Other attributes in this top-level object establish rules and restrictions
// that apply to all impressions being offered.
//
// There are also several subordinate objects that provide detailed data to
// potential buyers.
// Among these are the Site and App objects, which describe the type of
// published media in which the impression(s) appear.
// These objects are highly recommended, but only one applies to a given bid
// request depending on whether the media is browser-based web content or a
// non-browser application, respectively.
// See OpenRTB 2.5 Sec 3.2.1.
//go:generate easyjson $GOFILE
//easyjson:json
type BidRequest struct {
	// Attribute:
	//   id
	// Type:
	//   string; required
	// Description:
	//   Unique ID of the bid request, provided by the exchange.
	ID string `json:"id"`

	// Attribute:
	//   imp
	// Type:
	//   object array; required
	// Description:
	//   Array of Imp objects (Section 3.2.4) representing the
	//   impressions offered. At least 1 Imp object is required.
	Impressions []*Impression `json:"imp"`

	// Attribute:
	//   site
	// Type:
	//   object; recommended
	// Description:
	//    Details via a Site object (Section 3.2.13) about the publisher’s
	//    website. Only applicable and recommended for websites.
	Site *Site `json:"site,omitempty"`

	// Attribute:
	//   app
	// Type:
	//   object; recommended
	// Description:
	//    Details via an App object (Section 3.2.14) about the publisher’s
	//    app (i.e., non-browser applications). Only applicable and
	//    recommended for apps.
	Application *Application `json:"app,omitempty"`

	// Attribute:
	//   device
	// Type:
	//   object; recommended
	// Description:
	//   Details via a Device object (Section 3.2.18) about the user’s
	//   device to which the impression will be delivered.
	Device *Device `json:"device,omitempty"`

	// Attribute:
	//   user
	// Type:
	//   object; recommended
	// Description:
	//    Details via a User object (Section 3.2.20) about the human
	//    user of the device; the advertising audience.
	User *User `json:"user,omitempty"`

	// Attribute:
	//   test
	// Type:
	//   integer; default 0
	// Description:
	//    Indicator of test mode in which auctions are not billable,
	//    where 0 = live mode, 1 = test mode.
	IsTestMode NumericBool `json:"test,omitempty"`

	// Attribute:
	//   at
	// Type:
	//   integer; default 2
	// Description:
	//    Auction type, where 1 = First Price, 2 = Second Price Plus.
	//    Exchange-specific auction types can be defined using values
	//    greater than 500.
	AuctionType AuctionType `json:"at"`

	// Attribute:
	//   tmax
	// Type:
	//   integer
	// Description:
	//    Maximum time in milliseconds the exchange allows for bids to
	//    be received including Internet latency to avoid timeout. This
	//    value supersedes any a priori guidance from the exchange.
	Timeout int `json:"tmax,omitempty"`

	// Attribute:
	//   wseat
	// Type:
	//   string array
	// Description:
	//   White list of buyer seats (e.g., advertisers, agencies) allowed
	//   to bid on this impression. IDs of seats and knowledge of the
	//   buyer’s customers to which they refer must be coordinated
	//   between bidders and the exchange a priori. At most, only one
	//   of wseat and bseat should be used in the same request.
	//   Omission of both implies no seat restrictions.
	WhitelistedSeats []string `json:"wseat,omitempty"`

	// Attribute:
	//   bseat
	// Type:
	//   string array
	// Description:
	//   Block list of buyer seats (e.g., advertisers, agencies) restricted
	//   from bidding on this impression. IDs of seats and knowledge
	//   of the buyer’s customers to which they refer must be
	//   coordinated between bidders and the exchange a priori. At
	//   most, only one of wseat and bseat should be used in the
	//   same request. Omission of both implies no seat restrictions.
	BlocklistedSeats []string `json:"bseat,omitempty"`

	// Attribute:
	//   allimps
	// Type:
	//   integer; default 0
	// Description:
	//   Flag to indicate if Exchange can verify that the impressions
	//   offered represent all of the impressions available in context
	//   (e.g., all on the web page, all video spots such as pre/mid/post
	//   roll) to support road-blocking. 0 = no or unknown, 1 = yes, the
	//   impressions offered represent all that are available.
	HasAllImpressions NumericBool `json:"allimps,omitempty"`

	// Attribute:
	//   cur
	// Type:
	//   string array
	// Description:
	//   Array of allowed currencies for bids on this bid request using
	//   ISO-4217 alpha codes. Recommended only if the exchange
	//   accepts multiple currencies.
	Currencies []Currency `json:"cur,omitempty"`

	// Attribute:
	//   wlang
	// Type:
	//   string array
	// Description:
	//   White list of languages for creatives using ISO-639-1-alpha-2.
	//   Omission implies no specific restrictions, but buyers would be
	//   advised to consider language attribute in the Device and/or
	//   Content objects if available.
	WhitelistLanguages []string `json:"wlang,omitempty"`

	// Attribute:
	//   bcat
	// Type:
	//   string array
	// Description:
	//   Blocked advertiser categories using the IAB content
	//   categories. Refer to List 5.1.
	BlockedCategories []Category `json:"bcat,omitempty"`

	// Attribute:
	//   badv
	// Type:
	//   string array
	// Description:
	//   Block list of advertisers by their domains (e.g., “ford.com”).
	BlockedAdvertisers []string `json:"badv,omitempty"`

	// Attribute:
	//   bapp
	// Type:
	//   string array
	// Description:
	//   Block list of applications by their platform-specific exchange-independent
	//   application identifiers. On Android, these should
	//   be bundle or package names (e.g., com.foo.mygame). On iOS,
	//   these are numeric IDs.
	BlockedAdvertisersByMarketID []string `json:"bapp,omitempty"`

	// Attribute:
	//   source
	// Type:
	//   object
	// Description:
	//   A Sorce object (Section 3.2.2) that provides data about the
	//   inventory source and which entity makes the final decision.
	Regulation *Regulation `json:"regs,omitempty"`

	// Attribute:
	//   regs
	// Type:
	//   object
	// Description:
	//   A Regs object (Section 3.2.3) that specifies any industry, legal,
	//   or governmental regulations in force for this request.
	Source *Source `json:"source,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Extension interface{} `json:"ext,omitempty"`
}

// Validate method checks to see if the BidRequest object contains required and well-formatted data
// and returns a corresponding error when the check fails.
func (r *BidRequest) Validate() error {
	if r.ID == "" {
		return ErrInvalidBidRequestID
	}
	if r.Impressions == nil || len(r.Impressions) == 0 {
		return ErrInvalidBidRequestImpressions
	}

	if len(r.WhitelistedSeats) > 0 && len(r.BlocklistedSeats) > 0 {
		return ErrInvalidBidRequestSeats
	}

	if r.Application != nil && r.Site != nil {
		return ErrBidRequestHasBothAppAndSite
	}

	return nil
}

// String method returns a human-readable representation of the bid request object also suitable
// for logging with %s, or %v.
func (r *BidRequest) String() string {
	return fmt.Sprintf("[%s;%d]", r.ID, len(r.Impressions))
}

// Copy returns a pointer to a copy of the BidRequest object.
func (r *BidRequest) Copy() *BidRequest {
	if r == nil {
		return nil
	}

	brCopy := *r

	if r.Impressions != nil {
		brCopy.Impressions = []*Impression{}
		for i := range r.Impressions {
			brCopy.Impressions = append(brCopy.Impressions, r.Impressions[i].Copy())
		}
	}
	brCopy.Site = r.Site.Copy()
	brCopy.Application = r.Application.Copy()
	brCopy.Device = r.Device.Copy()
	brCopy.User = r.User.Copy()
	brCopy.WhitelistedSeats = util.DeepCopyStrSlice(r.WhitelistedSeats)
	brCopy.BlocklistedSeats = util.DeepCopyStrSlice(r.BlocklistedSeats)

	if r.Currencies != nil {
		brCopy.Currencies = make([]Currency, len(r.Currencies))
		copy(brCopy.Currencies, r.Currencies)
	}

	brCopy.WhitelistLanguages = util.DeepCopyStrSlice(r.WhitelistLanguages)

	if r.BlockedCategories != nil {
		brCopy.BlockedCategories = make([]Category, len(r.BlockedCategories))
		copy(brCopy.BlockedCategories, r.BlockedCategories)
	}

	brCopy.BlockedAdvertisers = util.DeepCopyStrSlice(r.BlockedAdvertisers)
	brCopy.BlockedAdvertisersByMarketID = util.DeepCopyStrSlice(r.BlockedAdvertisersByMarketID)
	brCopy.Regulation = r.Regulation.Copy()
	brCopy.Source = r.Source.Copy()
	brCopy.Extension = util.DeepCopyCopiable(r.Extension)
	return &brCopy
}
