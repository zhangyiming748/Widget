package getFileType

import (
	"allTools/util/log"
	"net/http"
	"os"
)

func Detect(fp string) string {
	// Open File
	f, err := os.Open(fp)
	if err != nil {
		log.Info.Println(err)
	}
	defer f.Close()
	// Get the content
	contentType, err := getFileContentType(f)
	if err != nil {
		log.Info.Println(err)
	}
	return contentType
}
func getFileContentType(out *os.File) (string, error) {
	// 只需要前 512 个字节就可以了
	buffer := make([]byte, 512)
	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}
	contentType := http.DetectContentType(buffer)
	return contentType, nil
}
