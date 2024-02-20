package core

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	. "komeet/config"
	"os"
	"path"
	"sync"
	"time"
)

var once sync.Once

var log zerolog.Logger

func Logger() zerolog.Logger {
	once.Do(func() {
		zerolog.TimeFieldFormat = time.RFC3339
		mw := io.MultiWriter(
			zerolog.ConsoleWriter{
				Out:        os.Stderr,
				TimeFormat: time.RFC3339,
			},
			newLogRotation(),
		)

		log = zerolog.New(mw).Level(zerolog.InfoLevel).With().Timestamp().Logger()
	})

	return log
}

func LoggerMiddleware(c *gin.Context) {
	start := time.Now()
	l := Logger()

	defer func() {
		l.Info().
			Str("method", c.Request.Method).
			Str("url", c.Request.URL.String()).
			Int("status", c.Writer.Status()).
			Str("user_agent", c.Request.UserAgent()).
			Dur("elapsed_ms", time.Since(start)).
			Msg("Incoming Request")
	}()

	c.Next()
}

func newLogRotation() io.Writer {
	return &lumberjack.Logger{
		Filename:   path.Join("logs/app.log"),
		MaxBackups: Config.Logs.MaxBackups,
		MaxSize:    Config.Logs.MaxSize,
		MaxAge:     Config.Logs.MaxDays,
		Compress:   Config.Logs.Compress,
	}
}
