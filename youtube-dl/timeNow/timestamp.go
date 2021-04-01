package timeNow

import "time"

func DateNowFormatStr() string {
	tm := time.Now()
	return tm.Format("2006-01-02 15:04:05")
}
