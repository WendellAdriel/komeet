package core

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path"
	"sync"
	"time"
)

var once sync.Once

var log zerolog.Logger

func (a *Application) Logger() zerolog.Logger {
	once.Do(func() {
		zerolog.TimeFieldFormat = time.RFC3339
		mw := io.MultiWriter(
			zerolog.ConsoleWriter{
				Out:        os.Stderr,
				TimeFormat: time.RFC3339,
			},
			newLogRotation(a),
		)

		log = zerolog.New(mw).Level(zerolog.InfoLevel).With().Timestamp().Logger()
	})

	return log
}

func loggerMiddleware(c *gin.Context) {
	start := time.Now()

	defer func() {
		log.Info().
			Str("method", c.Request.Method).
			Str("url", c.Request.URL.String()).
			Int("status", c.Writer.Status()).
			Str("user_agent", c.Request.UserAgent()).
			Dur("elapsed_ms", time.Since(start)).
			Msg("Incoming Request")
	}()

	c.Next()
}

func newLogRotation(a *Application) io.Writer {
	return &lumberjack.Logger{
		Filename:   path.Join("logs/app.log"),
		MaxBackups: a.Config.Logs.MaxBackups,
		MaxSize:    a.Config.Logs.MaxSize,
		MaxAge:     a.Config.Logs.MaxDays,
		Compress:   a.Config.Logs.Compress,
	}
}
