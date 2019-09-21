package upload

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"tqgin/conf"
	"tqgin/pkg/file"
	"tqgin/pkg/util"
)

func GetImageFullUrl(name string) string {
	return config.Tqconfig.String(config.PrefixUrl) + "/" + GetImagePath() + name
}

func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)
	return fileName + ext
}

func GetImagePath() string {
	return config.Tqconfig.String(config.ImageSavePath)
}

func GetImageFullPath() string {
	return config.Tqconfig.String(config.RuntimeRootPath) + GetImagePath()
}

func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)

	arraystring := config.Tqconfig.String(config.ImageAllowExts)

	for _, allowExt := range strings.Split(arraystring, ",") {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		return false
	}

	conSize, err := config.Tqconfig.Int(config.ImageMaxSize)
	if err != nil {
		log.Println(err)
		return false
	}

	log.Println("CheckImageSize size %v:%v", size, conSize)
	return size <= conSize
}

func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMKDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMKDir err :%v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %v", src)
	}
	return nil

}
