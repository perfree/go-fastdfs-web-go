package commons

import (
	"fmt"
	"os"
)

// FormatFileSize 文件大小格式化
func FormatFileSize(fileSize float64) (size string) {
	if fileSize < 1024 {
		//return strconv.FormatInt(fileSize, 10) + "B"
		return fmt.Sprintf("%.2fB", fileSize/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", fileSize/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", fileSize/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", fileSize/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fTB", fileSize/float64(1024*1024*1024*1024))
	} else { //if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2fEB", fileSize/float64(1024*1024*1024*1024*1024))
	}
}

// FormatFileSizeToMb 文件大小转Mb
func FormatFileSizeToMb(fileSize float64) (size string) {
	return fmt.Sprintf("%.2f", fileSize/float64(1024*1024))
}



// IsDirExists 目录是否存在
func IsDirExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}