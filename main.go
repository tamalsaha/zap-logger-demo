package main

import (
	"go.uber.org/zap/zapcore"
	"time"

	"go.uber.org/zap"
)

func main() {
	// Using zap's preset constructors is the simplest way to get a feel for the
	// package, but they don't allow much customization.

	zap.NewExample()
	zap.NewDevelopment()

	// NewConsoleEncoder

	//core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg.EncoderConfig), os.Stdout, zap.DebugLevel)
	//logger := zap.New(core)


	//zap.NewProduction()
	zap.NewExample()

	encoderCfg := zap.NewProductionConfig()
	encoderCfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	encoderCfg.Encoding = "console"
	encoderCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, _ := encoderCfg.Build() // zap.NewDevelopment() // or NewProduction, or NewDevelopment
	defer logger.Sync()

	const url = "http://example.com"

	// In most circumstances, use the SugaredLogger. It's 4-10x faster than most
	// other structured logging packages and has a familiar, loosely-typed API.
	sugar := logger.Sugar()
	sugar.Infow("Failed to fetch URL.",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)

	// In the unusual situations where every microsecond matters, use the
	// Logger. It's even faster than the SugaredLogger, but only supports
	// structured logging.
	logger.Debug("Failed to fetch URL.",
		// Structured context as strongly typed fields.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}
