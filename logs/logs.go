package logs

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger
var SugarLogger *zap.SugaredLogger

func InitLogger() {
	encoder := getEncoder()
	core1 := zapcore.NewCore(encoder, getLogWriter(), zapcore.DebugLevel)
	core2 := zapcore.NewCore(encoder, getLogErrWriter(), zapcore.ErrorLevel)

	core := zapcore.NewTee(core1, core2)
	Logger = zap.New(core)
	SugarLogger = Logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	workdir, _ := os.Getwd()
	file, _ := os.Create(workdir + "/logs/study_system.log")
	return zapcore.AddSync(file)
}
func getLogErrWriter() zapcore.WriteSyncer {
	workdir, _ := os.Getwd()
	file, _ := os.Create(workdir + "/logs/study_system.err.log")
	return zapcore.AddSync(file)
}
