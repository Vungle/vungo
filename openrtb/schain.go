package openrtb

// This object represents both the links in the supply chain as well as an indicator whether or not the supply chain is complete.
// The SupplyChain object should be included in the BidRequest.Source.ext.schain attribute in OpenRTB 2.5 or later. For OpenRTB 2.4 or prior, the BidRequest.ext.schain attribute should be used..
type SChain struct {

	// Attribute:
	//   complete
	// Type:
	//   Integer; required
	// Description:
	//   Flag indicating whether the chain contains all nodes involved in the transaction leading back to the owner of the site, app or other medium of the inventory, where
	//   0 = no, 1 = yes.
	Complete int8 `json:"complete"`

	// Attribute:
	//   nodes
	// Type:
	//   string; required
	// Description:
	//   Array of SupplyChainNode objects in the order of the chain. In a complete supply chain, the first node represents the initial advertising system and seller ID involved in the transaction, i.e. the owner of the site, app, or other medium. In an incomplete supply chain, it represents the first known node. The last node represents the entity sending this bid request.
	Nodes []SupplyChainNode `json:"nodes"`

	// Attribute:
	//   ver
	// Type:
	//   string; required
	// Description:
	//   Version of the supply chain specification in use, in the format of “major.minor”. For example, for version 1.0 of the spec, use the string “1.0”.
	Ver string `json:"ver,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext interface{} `json:"ext,omitempty"`
}

// Copy returns a pointer to a copy of the SChain object.
func (sc *SChain) Copy() *SChain {
	scCopy := *sc
	scCopy.Nodes = make([]SupplyChainNode, len(sc.Nodes))
	if sc.Nodes != nil {
		for _, node := range sc.Nodes {
			nodeCopy := node.Copy()
			scCopy.Nodes = append(scCopy.Nodes, *nodeCopy)
		}
	}
	// extension copying has to be done by the user of this package manually.
	scCopy.Ext = nil
	return &scCopy
}

// This object is associated with a SupplyChain object as an array of nodes.
// These nodes define the identity of an entity participating in the supply chain of a bid request.
type SupplyChainNode struct {

	// Attribute:
	//   asi
	// Type:
	//   String; required
	// Description:
	//   The canonical domain name of the SSP, Exchange, Header Wrapper, etc system that bidders connect to.
	//   This may be the operational domain of the system, if that is different than the parent corporate domain, to facilitate WHOIS and reverse IP lookups to establish clear ownership of the delegate system.
	//   This should be the same value as used to identify sellers in an ads.txt file if one exists.
	ASI int8 `json:"asi"`

	// Attribute:
	//   sid
	// Type:
	//   string; required
	// Description:
	//   The identifier associated with the seller or reseller account within the advertising system.
	//   This must contain the same value used in transactions (i.e. OpenRTB bid requests) in the field specified by the SSP/exchange.
	//   Typically, in OpenRTB, this is publisher.id.
	//   For OpenDirect it is typically the publisher’s organization ID.
	//   Should be limited to 64 characters in length.
	SID string `json:"sid"`

	// Attribute:
	//   rid
	// Type:
	//   string; optional
	// Description:
	//   The OpenRTB RequestId of the request as issued by this seller.
	RID string `json:"rid,omitempty"`

	// Attribute:
	//   name
	// Type:
	//   string; optional
	// Description:
	//   The name of the company (the legal entity) that is paid for inventory transacted under the given seller_id.
	//   This value is optional and should NOT be included if it exists in the advertising system’s sellers.json file.
	Name string `json:"name,omitempty"`

	// Attribute:
	//   domain
	// Type:
	//   string; optional
	// Description:
	//  The business domain name of the entity represented by this node. This value is optional and should NOT be included if it exists in the advertising system’s sellers.json file.
	Domain string `json:"domain,omitempty"`

	// Attribute:
	//   hp
	// Type:
	//   Integer; required
	// Description:
	//   Indicates whether this node will be involved in the flow of payment for the inventory. When set to 1, the advertising system in the asi field pays the seller in the sid field, who is responsible for paying the previous node in the chain. When set to 0, this node is not involved in the flow of payment for the inventory. For version 1.0 of SupplyChain, this property should always be 1. It is explicitly required to be included as it is expected that future versions of the specification will introduce non-payment handling nodes. Implementers should ensure that they support this field and propagate it onwards when constructing SupplyChain objects in bid requests sent to a downstream advertising system.
	HP int8 `json:"hp"`

	// Attribute:
	//   ext
	// Type:
	//   object; optional
	// Description:
	//   Placeholder for advertising-system specific extensions to this object.
	Ext interface{} `json:"ext,omitempty"`
}

// Copy returns a pointer to a copy of the SChain object.
func (scn *SupplyChainNode) Copy() *SupplyChainNode {
	scnCopy := *scn
	// extension copying has to be done by the user of this package manually.
	scnCopy.Ext = nil
	return &scnCopy
}
