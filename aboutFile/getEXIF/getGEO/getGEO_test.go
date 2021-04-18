package getGEO

import "testing"

func TestEXIF2GEO(t *testing.T) {
	ret := EXIF2GEO("/Users/zen/Github/Widget/aboutFile/getEXIF/IMG4.jpg")
	t.Logf("ret = %s", ret)
}
