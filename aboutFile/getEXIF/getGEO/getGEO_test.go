package getGEO

import "testing"

func TestEXIF2GEO(t *testing.T) {
	ret := EXIF2GEO("/Users/zen/Github/Widget/aboutFile/getEXIF/IMGX.jpg")
	t.Logf("ret = %s", ret)
}
