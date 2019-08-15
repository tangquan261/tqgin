package tqlog

import (
	"time"

	"github.com/keepeye/logrus-filename"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var TQSysLog *logrus.Logger
var TQRequest *logrus.Logger

func ConfigLog() {
	TQSysLog = NewSysLogger()
	TQRequest = NewRequestLogger()
}

func NewDBLogger() *logrus.Logger {

	path := "./dblog/dbinfo.log"

	writer, _ := rotatelogs.New(
		path+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(3600*24*30)*time.Second),    //文件最大保存时间
		rotatelogs.WithRotationTime(time.Duration(3600*24)*time.Second), //日志切割时间间隔
	)

	DBLog := logrus.New()
	DBLog.SetLevel(logrus.DebugLevel)
	filenameHook := filename.NewHook()
	filenameHook.Field = "line"
	DBLog.AddHook(filenameHook)
	DBLog.Hooks.Add(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  writer,
			logrus.ErrorLevel: writer,
			logrus.DebugLevel: writer,
		},
		&logrus.JSONFormatter{},
	))

	return DBLog
}

func NewSysLogger() *logrus.Logger {

	path := "./dblog/syslog.log"

	writer, _ := rotatelogs.New(
		path+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(3600*24*30)*time.Second),    //文件最大保存时间
		rotatelogs.WithRotationTime(time.Duration(3600*24)*time.Second), //日志切割时间间隔
	)

	patherror := "./dblog/errorlog.log"
	errorInfo, _ := rotatelogs.New(
		patherror+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(patherror),
		rotatelogs.WithMaxAge(time.Duration(3600*24*30)*time.Second),    //文件最大保存时间
		rotatelogs.WithRotationTime(time.Duration(3600*24)*time.Second), //日志切割时间间隔
	)

	pathpanic := "./dblog/paniclog.log"
	panicInfo, _ := rotatelogs.New(
		pathpanic+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(pathpanic),
		rotatelogs.WithMaxAge(time.Duration(3600*24*30)*time.Second),    //文件最大保存时间
		rotatelogs.WithRotationTime(time.Duration(3600*24)*time.Second), //日志切割时间间隔
	)

	SysLog := logrus.New()

	SysLog.SetLevel(logrus.DebugLevel)
	filenameHook := filename.NewHook()
	filenameHook.Field = "line"
	SysLog.AddHook(filenameHook)
	SysLog.Hooks.Add(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.PanicLevel: panicInfo, //log之后会panic()
			logrus.FatalLevel: panicInfo, //log之后会调用os.Exit(1)
			logrus.ErrorLevel: errorInfo, //Something failed but I'm not quitting.
			logrus.WarnLevel:  errorInfo, //警告日志
			logrus.InfoLevel:  writer,    //记录日志
			logrus.DebugLevel: writer,
		},
		&logrus.JSONFormatter{},
	))

	return SysLog
}

func NewRequestLogger() *logrus.Logger {

	path := "./dblog/request.log"

	requestInfo, _ := rotatelogs.New(
		path+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(3600*24*30)*time.Second),    //文件最大保存时间
		rotatelogs.WithRotationTime(time.Duration(3600*24)*time.Second), //日志切割时间间隔
	)

	pathreponse := "./dblog/response.log"
	reponseInfo, _ := rotatelogs.New(
		pathreponse+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(pathreponse),
		rotatelogs.WithMaxAge(time.Duration(3600*24*30)*time.Second),    //文件最大保存时间
		rotatelogs.WithRotationTime(time.Duration(3600*24)*time.Second), //日志切割时间间隔
	)

	Request := logrus.New()
	Request.SetLevel(logrus.DebugLevel)
	filenameHook := filename.NewHook()
	filenameHook.Field = "line"
	Request.AddHook(filenameHook)
	Request.Hooks.Add(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  requestInfo, //请求日志
			logrus.DebugLevel: reponseInfo, //返回日志
		},
		&logrus.JSONFormatter{},
	))

	return Request
}
