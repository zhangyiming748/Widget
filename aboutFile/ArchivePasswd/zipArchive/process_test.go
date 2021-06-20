package zipArchive

import (
	"testing"
)

func TestUnzip(t *testing.T) {
	fp := "./nier.zip"
	right := "passwd"
	worng := "password"
	UnZip(fp, worng)
	UnZip(fp, right)

}
