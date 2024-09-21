package image

import (
	"github.com/axis1114/blog/global"
	"github.com/axis1114/blog/models/res"
	"github.com/axis1114/blog/service/image_ser"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
		res.FailWithMessage("图片上传失败", c)
		return
	}

	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			res.FailWithMessage("图片上传失败", c)
			return
		}
	}

	var resList []image_ser.FileUploadResponse

	for _, file := range fileList {
		serviceRes := image_ser.ImageUploadService(file)
		if !serviceRes.IsSuccess {
			resList = append(resList, serviceRes)
			continue
		}
		err = c.SaveUploadedFile(file, serviceRes.FileName)
		if err != nil {
			global.Log.Error("save file err:", zap.Error(err))
			serviceRes.Msg = err.Error()
			serviceRes.IsSuccess = false
			resList = append(resList, serviceRes)
			continue
		}
		resList = append(resList, serviceRes)
	}
	res.OkWithData(resList, c)
}
