package boot

import (
	"fmt"
	"os"
	"strings"

	"github.com/bayik4/boilerplate-golang/internal/model"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func LoggerInit(cfg model.ConfigLogger) (*zap.Logger, error) {
	level := zapcore.InfoLevel

	if err := level.UnmarshalText([]byte(strings.ToLower(cfg.Level))); err != nil {
		return nil, err
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.LevelKey = "level"
	encoderConfig.MessageKey = "message"
	encoderConfig.CallerKey = "caller"
	encoderConfig.EncodeDuration = zapcore.MillisDurationEncoder

	var encoder zapcore.Encoder

	switch strings.ToLower(cfg.Format) {
	case "console":
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(
			"2006-01-02 15:04:05",
		)
		encoder = zapcore.NewConsoleEncoder(encoderConfig)

	case "json":
		encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
		encoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
		encoder = zapcore.NewJSONEncoder(encoderConfig)

	default:
		return nil, fmt.Errorf(
			"unsupported log format: %s",
			cfg.Format,
		)
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(os.Stdout),
		level,
	)

	logger := zap.New(
		core,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
		zap.Fields(
			zap.String("service.name", cfg.ServiceName),
			zap.String("environment", cfg.Environment),
		),
	)

	return logger, nil
}
