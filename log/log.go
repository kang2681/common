package log

import (
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	logrus.Entry
}

const (
	B int64 = 1 << (10 * iota)
	KB
	MB
	GB
)

type LoggerConf struct {
	Level    string // 日志级别
	SaveFile bool   // 保存到文件
	Path     string // 日志路径，到文件名，Rotation 文件保存在同目录
	MaxSize  int64  // 文件最大大小
	MaxAge   int    // 日志保留天数
}

func Init(c LoggerConf) io.Writer {
	logrus.SetLevel(getLogLevel(c.Level))
	logrus.SetReportCaller(true)
	writer := GetWriter(c)
	logrus.SetOutput(writer)
	return writer
}

func GetWriter(c LoggerConf) io.Writer {
	if !c.SaveFile {
		return os.Stdout
	}
	logf, err := rotatelogs.New(
		filepath.Join(c.Path+".%Y%m%d"),
		rotatelogs.WithLinkName(c.Path),
		rotatelogs.WithMaxAge(time.Duration(c.MaxAge)*24*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
		rotatelogs.WithRotationSize(c.MaxSize),
	)
	if err != nil {
		return os.Stdout
	}
	return logf
}

func getLogLevel(level string) logrus.Level {
	if lvl, err := logrus.ParseLevel(level); err == nil {
		return lvl
	}
	return logrus.InfoLevel
}

func NewWithUUID() *Logger {
	entry := logrus.WithField("uuid", uuid.New().String())
	return &Logger{
		*entry,
	}
}

func Trace(args ...interface{}) {
	logrus.Trace(args...)
}

func Debug(args ...interface{}) {
	logrus.Debug(args...)
}

func Print(args ...interface{}) {
	logrus.Print(args...)
}

func Info(args ...interface{}) {
	logrus.Info(args...)
}

func Warn(args ...interface{}) {
	logrus.Warn(args...)
}

func Warning(args ...interface{}) {
	logrus.Warning(args...)
}

func Error(args ...interface{}) {
	logrus.Error(args...)
}

func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

func Panic(args ...interface{}) {
	logrus.Panic(args...)
}

func Tracef(format string, args ...interface{}) {
	logrus.Tracef(format, args...)
}

func Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func Printf(format string, args ...interface{}) {
	logrus.Printf(format, args...)
}

func Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

func Warningf(format string, args ...interface{}) {
	logrus.Warningf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	logrus.Panicf(format, args...)
}

// Entry Println family functions

func Traceln(args ...interface{}) {
	logrus.Traceln(args...)
}

func Debugln(args ...interface{}) {
	logrus.Debugln(args...)
}

func Infoln(args ...interface{}) {
	logrus.Infoln(args...)
}

func Println(args ...interface{}) {
	logrus.Println(args...)
}

func Warnln(args ...interface{}) {
	logrus.Warnln(args...)
}

func Warningln(args ...interface{}) {
	logrus.Warningln(args...)
}

func Errorln(args ...interface{}) {
	logrus.Errorln(args...)
}

func Fatalln(args ...interface{}) {
	logrus.Fatalln(args...)
}

func Panicln(args ...interface{}) {
	logrus.Panicln(args...)
}
