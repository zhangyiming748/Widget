package getGEO

import (
	"fmt"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
	"os"
	"strconv"
	"strings"
)

func EXIF2GEO(fp string) string {
	defer func() {
		if err := recover(); err != nil {
			Debugln(err)
		}
	}()
	f, _ := os.Open(fp)
	x, _ := exif.Decode(f)
	Longitude, _ := x.Get(exif.GPSLongitude)
	Latitude, _ := x.Get(exif.GPSLatitude)
	long := convert(Longitude)
	lat := convert(Latitude)
	ret := fmt.Sprintf("%f,%v", long, lat)
	return ret
}
func convert(tag *tiff.Tag) float64 {
	step1 := fmt.Sprintf("%v", tag)
	step2 := strings.Replace(step1, "\"", "", -1)
	step3 := strings.Replace(step2, "[", "", -1)
	step4 := strings.Replace(step3, "]", "", -1)
	step5 := strings.Split(step4, ",")
	hs := strings.Split(step5[0], "/")[0]
	ms := strings.Split(step5[1], "/")[0]
	ss := strings.Split(step5[2], "/")[0]
	h, _ := strconv.Atoi(hs)
	m, _ := strconv.Atoi(ms)
	s, _ := strconv.Atoi(ss)
	ret := float64(h) + float64(m)/60 + float64(s)/1000000/3600
	return ret
}
