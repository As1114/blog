package image

import (
	"blog/global"
	"blog/models/res"
	"blog/service/image_ser"
	"github.com/gin-gonic/gin"
	"io/fs"
	"os"
)

func (i Image) ImageUpload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("文件不存在", c)
		return
	}

	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error("make the dir fail", err.Error())
			return
		}
	}

	var resList []image_ser.FileUploadResponse

	for _, file := range fileList {
		serviceRes := image_ser.ImageService{}.ImageUploadService(file)
		if !serviceRes.IsSuccess {
			resList = append(resList, serviceRes)
			continue
		}
		err = c.SaveUploadedFile(file, serviceRes.FileName)
		if err != nil {
			global.Log.Error("save file err:", err)
			serviceRes.Msg = err.Error()
			serviceRes.IsSuccess = false
			resList = append(resList, serviceRes)
			continue
		}
		resList = append(resList, serviceRes)
	}
	res.OkWithData(resList, c)
}
