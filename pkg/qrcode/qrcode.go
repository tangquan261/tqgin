package qrcode

import (
	"image/jpeg"

	"tqgin/pkg/file"
	"tqgin/pkg/util"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

type QrCode struct {
	URL    string
	Width  int
	Height int
	Ext    string
	Level  qr.ErrorCorrectionLevel
	Mode   qr.Encoding
}

const (
	EXT_JPG = ".jpg"
)

func NewQrCode(url string, width, height int, level qr.ErrorCorrectionLevel, mode qr.Encoding) *QrCode {
	return &QrCode{
		URL:    url,
		Width:  width,
		Height: height,
		Level:  level,
		Mode:   mode,
		Ext:    EXT_JPG,
	}
}

func GetQrCodeFileName(value string) string {
	return util.EncodeMD5(value)
}

func GetQrCodeFullPath() string {
	return "roomPath" + "stting.appSetting.qrCodeSavePath"
}

func GetQrCodeFullURL(name string) string {
	return "settting.appSetting.prefixurl" + "/" + "setting.AppSetting.QrCodeSavePath" + name
}

func (q *QrCode) Encode(path string) (string, string, error) {
	name := GetQrCodeFileName(q.URL) + q.Ext
	src := path + name

	if file.CheckNotExist(src) == true {

		code, err := qr.Encode(q.URL, q.Level, q.Mode)
		if err != nil {
			return "", "", err
		}

		code, err = barcode.Scale(code, q.Width, q.Height)
		if err != nil {
			return "", "", err
		}

		f, err := file.MustOpen(name, path)
		if err != nil {
			return "", "", err
		}
		defer f.Close()

		err = jpeg.Encode(f, code, nil)
		if err != nil {
			return "", "", err
		}
	}

	return name, path, nil
}
