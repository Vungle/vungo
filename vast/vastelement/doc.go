/*
Package entity implements IAB VAST 3.0 specification located at
https://www.iab.com/wp-content/uploads/2015/06/VASTv3_0.pdf
Documents and interact a mixture of VAST wrappers for VAST 3.0.

All vast element has already implemented Validate interface.
Validate() will validate the current vast element node, and will recursively validate its child nodes according to VAST
3.0 spec.

Example Usage

	// Turn raw XML to Go struct.
	xmlData, _ := ioutil.ReadFile("vast.xml")
	v := &vast.Vast{}
	xml.Unmarshal(xmlData, v)

	// using Validate to verify whether the Unmarshal object has errors.
	err := vast.Validate()

	// Using Go struct.
	fmt.Println(v.Ad)
*/
package vastelement
