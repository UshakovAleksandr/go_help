package logger

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LogMiddleware for app.
type LogMiddleware struct {
	*zap.Logger
}

// NewLoggerMiddleware returns handler logging for httprouter.
func NewLoggerMiddleware(logger *zap.Logger) func(next httprouter.Handle) httprouter.Handle {
	middleware := LogMiddleware{logger}
	return middleware.Handler
}

// Handler calculation custom logger.
func (l LogMiddleware) Handler(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		start := time.Now()

		rw := newResponseWriter(w)

		next(rw, r, ps)

		labels := []zapcore.Field{
			zap.Duration("duration", time.Since(start)),
			zap.Int("status_code", rw.status),
		}

		if rw.status == 500 || rw.status == 400 {
			l.Error("failed", labels...)
			return
		}

		l.Info("success", labels...)
	}
}

func newResponseWriter(w http.ResponseWriter) *responseWriterDelegator {
	return &responseWriterDelegator{ResponseWriter: w}
}

type responseWriterDelegator struct {
	http.ResponseWriter
	status      int
	written     int64
	wroteHeader bool
}

func (r *responseWriterDelegator) WriteHeader(code int) {
	r.status = code
	r.wroteHeader = true
	r.ResponseWriter.WriteHeader(code)
}

func (r *responseWriterDelegator) Write(b []byte) (int, error) {
	n, err := r.ResponseWriter.Write(b)
	r.written += int64(n)
	return n, err
}
