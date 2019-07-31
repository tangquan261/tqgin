package controllers

import (
	"log"
	"tqgin/common"

	"tqgin/pkg/errorcode"
	"tqgin/pkg/upload"

	"github.com/gin-gonic/gin"
)

type UploadController struct {
	tqgin.Controller
}

func (this *UploadController) RegisterRouter(router *gin.Engine) {
	router.POST("upload", this.uploadImage)
}

func (s *UploadController) uploadImage(con *gin.Context) {

	file, image, err := con.Request.FormFile("image")
	if err != nil {
		log.Println(err)
		tqgin.ResultFail(con, err)
		return
	}

	if nil == image {
		tqgin.Result(con, errorcode.ERROR_INVALID_PARAMS, nil, "")
		return
	}
	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImageFullPath()
	//savePath := upload.GetImagePath()
	src := fullPath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		tqgin.Result(con, errorcode.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, nil, "")
		return
	}

	err = upload.CheckImage(fullPath)
	if err != nil {
		log.Println(err)
		tqgin.Result(con, errorcode.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil, "")
		return
	}

	if err := con.SaveUploadedFile(image, src); err != nil {
		log.Println(err)
		tqgin.Result(con, errorcode.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil, "")
		return
	}

	tqgin.ResultOk(con, nil)

}
