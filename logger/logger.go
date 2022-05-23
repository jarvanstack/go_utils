package logger

import (
	"io/ioutil"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v2"
)

//参考这个文章
//https://studygolang.com/articles/32205?fr=sidebar
var (
	//全局唯一 logger 变量
	logger *zap.SugaredLogger
	//是否开启debug,用于logger.Debug()的时候加一层判断是否输出 debug
	isDebug bool = true
)

type Config struct {
	IsDebug bool `yaml:"isDebug"`
}

func init() {
	var err error
	var config = new(Config)

	//判断是否 isDebug 默认为 true
	filePath := "./config/config.yaml"
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		isDebug = true
	} else {
		err = yaml.Unmarshal(b, config)
		if err == nil && config != nil && !config.IsDebug {
			isDebug = false
		}
	}

	//初始化 logger
	encoderConfig := zap.NewProductionEncoderConfig()
	//这个是标准但是不好看的时间
	// encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 自定义时间输出格式
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("2006-01-02 15:04:05.000")
	}
	encoderConfig.EncodeTime = customTimeEncoder
	encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	//文件writeSyncer
	// fileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
	// 	Filename:   "./logs/app.log", //日志文件存放目录
	// 	MaxSize:    1024,             //文件大小限制,单位MB
	// 	MaxBackups: 5,                //最大保留日志文件数量
	// 	MaxAge:     30,               //日志文件保留天数
	// 	Compress:   false,            //是否压缩处理
	// })
	//控制台输出 zapcore.NewMultiWriteSyncer(fileWriteSyncer,zapcore.AddSync(os.Stdout))
	consoleWriteSyncer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
	core := zapcore.NewCore(encoder, consoleWriteSyncer, zapcore.DebugLevel)
	//AddCaller()为显示文件名和行号
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()
}

func Sugar() *zap.SugaredLogger {
	return logger
}
func Debug(args ...interface{}) {
	if isDebug {
		logger.Debug(args)
	}
}
func Info(args ...interface{}) {
	logger.Info(args)
}
func Error(args ...interface{}) {
	logger.Error(args)
}
func Panic(args ...interface{}) {
	logger.Panic(args)
}
func Fatal(args ...interface{}) {
	logger.Fatal(args)
}

func Debugf(template string, args ...interface{}) {
	if isDebug {
		logger.Debugf(template, args...)
	}
}
func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}
func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}
func Panicf(template string, args ...interface{}) {
	logger.Panicf(template, args...)
}
func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args...)
}
