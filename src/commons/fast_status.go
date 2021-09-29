package commons

import (
	"time"
)

func GetStatus(data map[string]interface{}) (map[string]interface{}, error) {
	// 结果集
	result := make(map[string]interface{})
	// 声明30天文件大小数据容器
	var dayFileSizeList []string
	var dayFileCountList []float64
	// 声明30天内日期容器
	var dayNumList []string

	diskInfo := data["Sys.DiskInfo"].(map[string]interface{})
	result["diskFreeSize"] = FormatFileSize(diskInfo["free"].(float64))
	result["diskTotalSize"] = FormatFileSize(diskInfo["total"].(float64))
	result["diskUsedSize"] = FormatFileSize(diskInfo["used"].(float64))
	result["inodesTotal"] = diskInfo["inodesTotal"].(float64)
	result["inodesUsed"] = diskInfo["inodesUsed"].(float64)
	result["inodesFree"] = diskInfo["inodesFree"].(float64)

	dayFileSize := 0.0
	dayFileCount := 0.0

	fileStats := data["Fs.FileStats"].([]interface{})
	for _, v := range fileStats {
		fileStat := v.(map[string]interface{})
		if fileStat["date"] == "all" {
			result["totalFileSize"] = FormatFileSize(fileStat["totalSize"].(float64))
			result["totalFileCount"] = fileStat["fileCount"].(float64)
		} else {
			var formatTime, _ = time.Parse("20060102", fileStat["date"].(string))
			t := time.Now()
			beforeMonth := t.AddDate(0, -1, 0)
			if formatTime.Unix() <= t.Unix() && formatTime.Unix() >= beforeMonth.Unix() {
				dayFileSize += fileStat["totalSize"].(float64)
				dayFileCount += fileStat["fileCount"].(float64)
				dayFileSizeList = append(dayFileSizeList, FormatFileSizeToMb(fileStat["totalSize"].(float64)))
				dayFileCountList = append(dayFileCountList, fileStat["fileCount"].(float64))
				dayNumList = append(dayNumList, fileStat["date"].(string)[6:len(fileStat["date"].(string))]+"日")
			}
		}

	}

	result["dayFileSize"] = FormatFileSize(dayFileSize)
	result["dayFileCount"] = dayFileCount
	result["dayFileSizeList"] = dayFileSizeList
	result["dayFileCountList"] = dayFileCountList
	result["dayNumList"] = dayNumList
	return result, nil
}
