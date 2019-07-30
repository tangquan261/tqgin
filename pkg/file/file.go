package file

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

//获取文件大小size
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

//获取文件扩展名
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

//判断文件是否存在
func CheckNotExist(str string) bool {
	_, err := os.Stat(str)
	return os.IsNotExist(err)
}

//判断文件的操作权限
func CheckPermission(str string) bool {
	_, err := os.Stat(str)
	return os.IsPermission(err)
}

//不存在则创建文件夹
func IsNotExistMKDir(str string) error {
	if notExist := CheckNotExist(str); notExist == true {
		return os.MkdirAll(str, os.ModePerm)
	}
	return nil
}

func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	return f, nil
}

//打开一个文件
func MustOpen(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.getwd err %v", err)
	}

	src := dir + "/" + filePath

	perm := CheckPermission(src)

	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = IsNotExistMKDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}
	return f, nil
}
