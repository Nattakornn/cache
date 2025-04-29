package logger

import (
	"os"

	"github.com/Nattakornn/cache/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func InitZapLogger(cfg config.ILogConfig) {

	logLevel := getZapLevel(cfg.Level())
	logColor := cfg.Color()
	logJson := cfg.Json()

	var logEncoder zapcore.Encoder
	logEncoderConfig := zap.NewProductionEncoderConfig()
	logEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	if logColor {
		logEncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	}

	if logJson {
		logEncoder = zapcore.NewJSONEncoder(logEncoderConfig)
	} else {
		logEncoder = zapcore.NewConsoleEncoder(logEncoderConfig)
	}

	core := zapcore.NewCore(
		logEncoder,
		os.Stdout,
		zap.NewAtomicLevelAt(logLevel),
	)

	Logger = zap.New(core, zap.AddCaller()).Sugar()
	Logger.Infof("init zap logger complete")
}

func SyncLogger() {
	Logger.Infof("flush logger")
	Logger.Sync()
}

func getZapLevel(level string) zapcore.Level {
	switch level {
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "debug":
		return zapcore.DebugLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
