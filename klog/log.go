package klog

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/google/uuid"
)

var default_log = new("info", os.Stdout)
var log *zap.Logger = default_log

func InitWithConfig(c *LogConfig) {
	if c == nil || !c.SaveFile {
		log = default_log
	} else {
		log = new(c.Level, &c.Logger)
	}
}

func Init(level string, ws ...io.Writer) {
	log = new(level, ws...)
}

func new(level string, ws ...io.Writer) *zap.Logger {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder   // 时间格式
	cfg.EncodeCaller = zapcore.ShortCallerEncoder // 短路径
	cfg.EncodeLevel = zapcore.CapitalLevelEncoder // 大写级别
	cfg.EncodeName = zapcore.FullNameEncoder      //

	// 日志级别
	logLevel, err := zapcore.ParseLevel(level)
	if err != nil {
		logLevel = zapcore.InfoLevel
	}

	cores := make([]zapcore.Core, 0, len(ws))
	for _, w := range ws {
		core := zapcore.NewCore(zapcore.NewJSONEncoder(cfg), zapcore.AddSync(w), logLevel)
		cores = append(cores, core)
	}
	return zap.New(zapcore.NewTee(cores...), zap.AddCaller(), zap.Development())
}

func Sync() error {
	if log != nil {
		return log.Sync()
	}
	return nil
}

func NewLog() *zap.Logger {
	u, _ := uuid.NewUUID()
	return log.With(zap.String("uuid", u.String()))
}

func NewSugarLog() *zap.SugaredLogger {
	u, _ := uuid.NewUUID()
	return log.Sugar().With("uuid", u.String())
}

type LogConfig struct {
	lumberjack.Logger
	Level    string
	SaveFile bool
}

var DefaultLoggerConfig = &LogConfig{
	Logger: lumberjack.Logger{
		Filename:   "./logs/trace.log",
		MaxSize:    500,
		MaxAge:     7,
		MaxBackups: 30,
		Compress:   true,
	},
	Level:    "info",
	SaveFile: true,
}

func AddFlag(c *LogConfig) {
	flag.StringVar(&c.Filename, "log.path", DefaultLoggerConfig.Filename, " Filename is the file to write logs to")
	flag.IntVar(&c.MaxSize, "log.max-size", DefaultLoggerConfig.MaxSize, "MaxSize is the maximum size in megabytes of the log file before it gets rotated. ")
	flag.IntVar(&c.MaxAge, "log.max-age", DefaultLoggerConfig.MaxAge, "MaxAge is the maximum number of days to retain old log files based on the timestamp encoded in their filename")
	flag.IntVar(&c.MaxBackups, "log.max-backup", DefaultLoggerConfig.MaxBackups, " MaxBackups is the maximum number of old log files to retain")
	flag.BoolVar(&c.Compress, "log.compress", DefaultLoggerConfig.Compress, " Compress determines if the rotated log files should be compressed using gzip.")
	flag.StringVar(&c.Level, "log.level", DefaultLoggerConfig.Level, "Log level. debug,	info, warn, error, dpanic, panic, fatal")
	flag.BoolVar(&c.SaveFile, "log.save-file", DefaultLoggerConfig.SaveFile, "Log message write to file")
}

func AddKinpin(c *LogConfig) {
	kingpin.Flag("log.path", "Filename is the file to write logs to").Default(DefaultLoggerConfig.Filename).StringVar(&c.Filename)
	kingpin.Flag("log.max-size", "MaxSize is the maximum size in megabytes of the log file before it gets rotated.").Default(fmt.Sprintf("%d", DefaultLoggerConfig.MaxSize)).IntVar(&c.MaxSize)
	kingpin.Flag("log.max-age", "MaxAge is the maximum number of days to retain old log files based on the timestamp encoded in their filename").Default(fmt.Sprintf("%d", DefaultLoggerConfig.MaxAge)).IntVar(&c.MaxAge)
	kingpin.Flag("log.max-backup", "MaxBackups is the maximum number of old log files to retain").Default(fmt.Sprintf("%d", DefaultLoggerConfig.MaxBackups)).IntVar(&c.MaxBackups)
	kingpin.Flag("log.compress", "Compress determines if the rotated log files should be compressed using gzip.").Default(fmt.Sprintf("%t", DefaultLoggerConfig.Compress)).BoolVar(&c.Compress)
	kingpin.Flag("log.level", "Log level. debug,	info, warn, error, dpanic, panic, fatal").Default(DefaultLoggerConfig.Level).StringVar(&c.Level)
	kingpin.Flag("log.save-file", "Log save to file").Default(fmt.Sprintf("%t", DefaultLoggerConfig.SaveFile)).BoolVar(&c.SaveFile)
}
