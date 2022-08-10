package main

import (
	"flag"
	"fmt"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var Logger *zap.Logger

func Init(level string, w *LogConfig) {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder   // 时间格式
	cfg.EncodeCaller = zapcore.ShortCallerEncoder // 全路径
	cfg.EncodeLevel = zapcore.CapitalLevelEncoder // 大写级别
	cfg.EncodeName = zapcore.FullNameEncoder      //

	// 日志级别
	logLevel, err := zapcore.ParseLevel(level)
	if err != nil {
		logLevel = zapcore.InfoLevel
	}

	fileCore := zapcore.NewCore(zapcore.NewJSONEncoder(cfg), zapcore.AddSync(w), logLevel)
	consoleCore := zapcore.NewCore(zapcore.NewJSONEncoder(cfg), zapcore.AddSync(w), logLevel)

	cores := make([]zapcore.Core, 0, 2)
	if !w.Stdout && !w.OutputFile {
		w.Stdout = true
	}
	if w.Stdout {
		cores = append(cores, fileCore, consoleCore)
	}
	if w.OutputFile {
		cores = append(cores, fileCore)
	}
	Logger = zap.New(zapcore.NewTee(cores...), zap.AddCaller(), zap.Development())
}

type LogConfig struct {
	lumberjack.Logger
	Level      string
	OutputFile bool
	Stdout     bool
}

var LoggerConfig = &LogConfig{}

var DefaultLoggerConfig = &LogConfig{
	Logger: lumberjack.Logger{
		Filename:   "./logs/trace.log",
		MaxSize:    500,
		MaxAge:     7,
		MaxBackups: 30,
		Compress:   true,
	},
	Level:      "info",
	OutputFile: true,
	Stdout:     false,
}

func AddFlag() {
	flag.StringVar(&LoggerConfig.Filename, "log.path", DefaultLoggerConfig.Filename, " Filename is the file to write logs to")
	flag.IntVar(&LoggerConfig.MaxSize, "log.max-size", DefaultLoggerConfig.MaxSize, "MaxSize is the maximum size in megabytes of the log file before it gets rotated. ")
	flag.IntVar(&LoggerConfig.MaxAge, "log.max-age", DefaultLoggerConfig.MaxAge, "MaxAge is the maximum number of days to retain old log files based on the timestamp encoded in their filename")
	flag.IntVar(&LoggerConfig.MaxBackups, "log.max-backup", DefaultLoggerConfig.MaxBackups, " MaxBackups is the maximum number of old log files to retain")
	flag.BoolVar(&LoggerConfig.Compress, "log.compress", DefaultLoggerConfig.Compress, " Compress determines if the rotated log files should be compressed using gzip.")
	flag.StringVar(&LoggerConfig.Level, "log.level", DefaultLoggerConfig.Level, "Log level. debug,	info, warn, error, dpanic, panic, fatal")
	flag.BoolVar(&LoggerConfig.OutputFile, "log.output-file", DefaultLoggerConfig.OutputFile, "Log message write to file")
	flag.BoolVar(&LoggerConfig.Stdout, "log.stdout", DefaultLoggerConfig.Stdout, "Log message write to stdout")
}

func AddKinpin() {
	kingpin.Flag("log.path", "Filename is the file to write logs to").Default(DefaultLoggerConfig.Filename).StringVar(&LoggerConfig.Filename)
	kingpin.Flag("log.max-size", "MaxSize is the maximum size in megabytes of the log file before it gets rotated.").Default(fmt.Sprintf("%d", DefaultLoggerConfig.MaxSize)).Int64Var(&LoggerConfig.MaxSize)
	kingpin.Flag("log.max-age", "MaxAge is the maximum number of days to retain old log files based on the timestamp encoded in their filename").Default(fmt.Sprintf("%d", DefaultLoggerConfig.MaxAge)).IntVar(&LoggerConfig.MaxAge)
	kingpin.Flag("log.max-backup", "MaxBackups is the maximum number of old log files to retain").Default(fmt.Sprintf("%d", DefaultLoggerConfig.MaxBackups)).IntVar(&LoggerConfig.MaxBackups)
	kingpin.Flag("log.compress", "Compress determines if the rotated log files should be compressed using gzip.").Default(fmt.Sprintf("%t", DefaultLoggerConfig.Compress)).BoolVar(&LoggerConfig.Compress)
	kingpin.Flag("log.level", "Log level. debug,	info, warn, error, dpanic, panic, fatal").Default(DefaultLoggerConfig.Level).StringVar(&LoggerConfig.Level)
	kingpin.Flag("log.output-file", "Log save to file").Default(fmt.Sprintf("%t", DefaultLoggerConfig.OutputFile)).BoolVar(&LoggerConfig.OutputFile)
	kingpin.Flag("log.stdout", "Log message write to stdout").Default(fmt.Sprintf("%t", DefaultLoggerConfig.Stdout)).BoolVar(&LoggerConfig.Stdout)
}
