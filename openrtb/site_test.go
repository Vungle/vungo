package openrtb_test

import (
	"testing"

	"github.com/Vungle/vungo/openrtb"
	"github.com/Vungle/vungo/openrtb/openrtbtest"
)

func TestSite_Fields(t *testing.T) {
	standardText := `
Attribute
Type
Description
id
string; recommended
Exchange-specific site ID.
name
string
Site name (may be aliased at the publisher’s request).
domain
string
Domain of the site (e.g., “mysite.foo.com”).
cat
string array
Array of IAB content categories of the site. Refer to List 5.1.
sectioncat
string array
Array of IAB content categories that describe the current section of the site. Refer to List 5.1.
pagecat
string array
Array of IAB content categories that describe the current page or view of the site. Refer to List 5.1.
page
string
URL of the page where the impression will be shown.
ref
string
Referrer URL that caused navigation to the current page.
search
string
Search string that caused navigation to the current page.
mobile
integer
Indicates if the site has been programmed to optimize layout when viewed on mobile devices, where 0 = no, 1 = yes.
privacypolicy
integer
Indicates if the site has a privacy policy, where 0 = no, 1 = yes.
publisher
object
Details about the Publisher (Section 3.2.15) of the site.
content
object
Details about the Content (Section 3.2.16) within the site.
keywords
string
Comma separated list of keywords about the site.
ext
object
Placeholder for exchange-specific extensions to OpenRTB.`
	if err := openrtbtest.VerifyStructFieldNameWithStandardText(
		(*openrtb.Site)(nil), standardText); err != "" {
		t.Error(err)
	}
}

func TestSite_Copy(t *testing.T) {
	site := openrtb.Site{}
	openrtbtest.FillWithNonNilValue(&site)
	siteCopyPtr := site.Copy()
	if err := openrtbtest.VerifyDeepCopy(&site, siteCopyPtr); err != nil {
		t.Errorf("Copy() should be deep copy\n%v\n", err)
	}
}
