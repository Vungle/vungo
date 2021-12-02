/*
Package vast implements IAB VAST 3.0 specification located at
http://www.iab.net/media/file/VASTv3.0.pdf. The vast package can also be used to process VAST 2.0
documents and interact a mixture of VAST wrappers between VAST 2.0 and 3.0 documents.

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
package vast2
