package routers

import (
	"github.com/gin-gonic/gin"
	"go-fastdfs-web-go/src/commons"
	"go-fastdfs-web-go/src/commons/httplib"
	"io/ioutil"
	"net/http"
	"os"
)

type FileRouter struct {
	BaseRouter
	commons.Response
}

// Index 文件列表
func (f *FileRouter)Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "file/file.tpl", nil)
}

// GetDirFile 获取目录/文件列表
func (f *FileRouter) GetDirFile(ctx *gin.Context) {
	dir := ctx.PostForm("dir")
	url,_ :=f.GetPeersUrl(ctx)
	result,_ := commons.GetDirOrFileList(f.GetShowUrl(ctx), url, dir)
	f.SuccessData(ctx, result)
}

// DeleteFile 删除文件
func (f *FileRouter)DeleteFile (ctx *gin.Context)  {
	md5 := ctx.PostForm("md5")
	url,_ :=f.GetPeersUrl(ctx)
	if commons.DeleteFile(url, md5) {
		f.Success(ctx)
		return
	}
	f.Error(ctx)
}

// Details 文件信息
func (f *FileRouter)Details (ctx *gin.Context)  {
	md5 := ctx.PostForm("md5")
	url,_ :=f.GetPeersUrl(ctx)
	result, err := commons.Details(url,f.GetShowUrl(ctx), md5)
	if err != nil {
		f.ErrorMsg(ctx, "获取文件信息失败")
		return
	}
	f.SuccessData(ctx, result)
}

// DownloadFile 下载
func (f *FileRouter) DownloadFile(ctx *gin.Context){
	name := ctx.PostForm("name")
	path := ctx.PostForm("path")
	peersUrl,_ :=f.GetPeersUrl(ctx)
	res, err :=http.Get(peersUrl + "/" + path + "/" + name)
	if err !=nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	ctx.Writer.Header().Add("Content-Type", "application/octet-stream")
	ctx.Writer.Header().Add("Content-Disposition", "attachment; filename=\""+name+"\"")
	_, _ = ctx.Writer.Write(content)
}

// UploadPage 上传页
func (f *FileRouter) UploadPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "file/upload.tpl", gin.H{
		"showAddress" : f.GetShowUrl(ctx),
	})
}

// Upload 文件上传
func (f *FileRouter)Upload(ctx *gin.Context)  {
	file, _ := ctx.FormFile("file")
	if !commons.IsDirExists("tmp") {
		err := os.Mkdir("tmp", 0777)
		if err != nil {
			f.ErrorMsg(ctx, "创建临时目录失败")
			return
		}
	}

	filePath := "tmp/" + file.Filename
	err := ctx.SaveUploadedFile(file, filePath)
	if err != nil {
		f.ErrorMsg(ctx, "保存临时文件失败")
		return
	}

	peersUrl,_ :=f.GetPeersUrl(ctx)
	scene := ctx.PostForm("scene")
	path := ctx.PostForm("path")

	var obj map[string]interface{}
	req := httplib.Post(peersUrl + commons.ApiUpload)
	req.PostFile("file", filePath)
	req.Param("output","json")
	req.Param("scene",scene)
	req.Param("path",path)
	err = req.ToJSON(&obj)
	if err != nil {
		f.Error(ctx)
		return
	}
	obj["url"] = f.GetShowUrlNotGroup(ctx) + obj["path"].(string)
	err = os.Remove(filePath)
	f.SuccessData(ctx, obj)
}
