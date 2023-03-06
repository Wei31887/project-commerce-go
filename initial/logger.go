package initial

import (
	"os"
	"path/filepath"
	"project/e-commerce/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogMaker struct {
	Config config.Config
}

// Logger : initialize the log and return zap logger
func LoadingLogger(config config.Config) (logger *zap.Logger) {
	maker := &LogMaker{
		Config: config,
	}
	core := maker.getLogCore()
	logger = zap.New(core)
	return logger
}

// 
func (l *LogMaker) getLogCore() zapcore.Core {
	return zapcore.NewCore(getEncoder(l.Config), getLogWriter(l.Config), getCoreLevel(l.Config))
}

func filePathName(config config.Config) string {
	return filepath.Join(config.Mylog.Path + "/" + config.Mylog.Name + ".log")
}

// indicate the location of log file 
func getLogWriter(config config.Config) zapcore.WriteSyncer {
	// Outputting log to console or file is depend on the config setting
	if config.Mylog.Model == "console" {
		return zapcore.AddSync(os.Stdout)
	} else {
		file, _ := os.Create(filePathName(config))
		return zapcore.AddSync(file)
	}
}

// 
func getEncoder(config config.Config) zapcore.Encoder {
	if config.Mylog.Format == "JSON" {
		return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
}

// 
func getCoreLevel(config config.Config) zapcore.Level {
	switch config.Mylog.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}




