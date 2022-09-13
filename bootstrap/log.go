package bootstrap

import (
	"jassue-gin/global"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// InitZapLogger Init zap logger and set to global log
func InitZapLogger() {
	logLevel := global.App.Config.Log.Level
	logFormat := global.App.Config.Log.Format

	// Set log level
	zapLogLevel := zapcore.DebugLevel
	switch logLevel {
	case "debug":
		zapLogLevel = zapcore.DebugLevel
	case "info":
		zapLogLevel = zapcore.InfoLevel
	case "warn":
		zapLogLevel = zapcore.WarnLevel
	case "error":
		zapLogLevel = zapcore.ErrorLevel
	case "fatal":
		zapLogLevel = zapcore.FatalLevel
	case "panic":
		zapLogLevel = zapcore.PanicLevel
	default:
		zapLogLevel = zapcore.DebugLevel
	}

	// Set log format
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006/01/02 15:04:05.999")
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// Set log output
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	if logFormat == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	// Set log rotate
	lumberJackLogger := &lumberjack.Logger{
		Filename:   global.App.Config.Log.RootDir + "/" + global.App.Config.Log.Filename,
		MaxSize:    global.App.Config.Log.MaxSize,
		MaxBackups: global.App.Config.Log.MaxBackups,
		MaxAge:     global.App.Config.Log.MaxAge,
		Compress:   global.App.Config.Log.Compress,
		LocalTime:  true,
	}

	// Init zap
	fileSyncer := zapcore.AddSync(lumberJackLogger)
	consoleSyncer := zapcore.AddSync(os.Stdout)
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(fileSyncer, consoleSyncer), zapLogLevel)
	zapLogger := zap.New(core, zap.AddCaller())

	// Zap to log
	zap.ReplaceGlobals(zapLogger)
	defer zapLogger.Sync()
	ZapToLog()
	OutConfig(zapLogger)
	zapLogger.Info("Init zap logger success")
}
