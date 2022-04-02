package logger

import (
	"webapp/settings"

	"github.com/natefinch/lumberjack"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Init(c *settings.App) (err error) {
	cl := &lumberjack.Logger{
		Filename:   c.Log.File,
		MaxSize:    c.MaxSize,
		MaxAge:     c.MaxAge,
		MaxBackups: c.MaxBackUp,
		Compress:   false,
	}
	l := new(zapcore.Level)
	err = l.UnmarshalText([]byte(c.Level))
	core := zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), zapcore.AddSync(cl), l)
	logger := zap.New(core)
	zap.ReplaceGlobals(logger)
	return
}
