/*
Package vast2 implements IAB VAST 2.0 specification located at
https://iabtechlab.com/wp-content/uploads/2016/04/VAST-2_0-FINAL.pdf.
Documents and interact a mixture of VAST wrappers for VAST 2.0.

All vast element has already implemented Validate interface.
Validate() will validate the current vast element node, and will recursively validate its child nodes according to VAST
2.0 spec.

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
