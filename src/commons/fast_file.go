package commons

import (
	"encoding/json"
	"net/url"
	"strings"
	"time"
)

// GetDirOrFileList 获取文件列表
func GetDirOrFileList(showUrl string, serverAddress string, dir string) ([]map[string]interface{}, error){
	var fileList []map[string]interface{}
	postValue := url.Values{}
	if dir != "" {
		postValue.Set("dir", dir)
	}
	var httpUtil HttpUtil
	result, err := httpUtil.PostForm(serverAddress + ApiListDir, postValue)
	if err != nil || result == ""{
		return fileList, err
	}

	var resultMap map[string]interface{}
	err = json.Unmarshal([]byte(result), &resultMap)
	if err != nil {
		return fileList, err
	}

	if resultMap["status"] == ApiStatusSuccess && resultMap["data"] != nil && len( resultMap["data"].([]interface{})) > 0 {
		data := resultMap["data"].([]interface{})
		for _,v := range data {
			value := v.(map[string]interface{})
			if value["name"] == "_big" {
				continue
			}

			fileResult := make(map[string]interface{})
			fileResult["md5"] = value["md5"]
			fileResult["path"] = value["path"]
			fileResult["name"] = value["name"]
			fileResult["is_dir"] = value["is_dir"].(bool)
			fileResult["peerAddr"] = showUrl
			if value["is_dir"].(bool) {
				fileResult["size"] = 0
			} else {
				fileResult["size"] = FormatFileSize(value["size"].(float64))
			}
			mTime := time.Unix(int64(value["mtime"].(float64)), 0).Format("2006-01-02 15:04:05")
			fileResult["mTime"] = mTime
			fileList = append(fileList, fileResult)
		}
	}
	return fileList, err
}

// DeleteFile 删除文件
func DeleteFile(peersUrl string, md5 string) bool {
	postValue := url.Values{"md5": {md5}}
	var httpUtil HttpUtil
	result, err := httpUtil.PostForm(peersUrl + ApiDelete, postValue)
	if err != nil{
		return false
	}

	var resultMap map[string]interface{}
	err = json.Unmarshal([]byte(result), &resultMap)
	if err != nil {
		return false
	}

	if resultMap["status"] == ApiStatusSuccess{
		return true
	}
	return false
}

// Details 文件信息
func Details(peersUrl string,showUrl string, md5 string) (map[string]interface{}, error) {
	postValue := url.Values{"md5": {md5}}
	fileResult := make(map[string]interface{})
	var httpUtil HttpUtil
	result, err := httpUtil.PostForm(peersUrl + ApiGetFileInfo, postValue)
	if err != nil{
		return nil, err
	}

	var resultMap map[string]interface{}
	err = json.Unmarshal([]byte(result), &resultMap)
	if err != nil {
		return nil, err
	}

	if resultMap["status"] == ApiStatusSuccess{
		resultMap = resultMap["data"].(map[string]interface{})
		fileResult["url"] = showUrl + "/" + strings.Replace(resultMap["path"].(string), "files/", "", 1) + "/" + resultMap["name"].(string)
		fileResult["path"] = resultMap["path"]
		fileResult["size"] = FormatFileSize(resultMap["size"].(float64))
		fileResult["name"] = resultMap["name"]
		fileResult["md5"] = resultMap["md5"]
		fileResult["scene"] = resultMap["scene"]
		fileResult["timeStamp"] = time.Unix(int64(resultMap["timeStamp"].(float64)), 0).Format("2006-01-02 15:04:05")
		return fileResult, nil
	}
	return nil, err
}