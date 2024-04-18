package middleware

import (
	"go-projects/social-media-profile-management/pkg/log"
	"net/http"
	"time"
)

type ResponseWriter struct {
	http.ResponseWriter
	status int
}

var logger = log.GetLogger()

func LoggerMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Start the timer
		start := time.Now()

		// Wrap the response writer, so we can capture the status code
		wrappedWriter := NewResponseWriter(w)

		// Call the next handler. which can be another middleware in the chain or the final handler
		next.ServeHTTP(wrappedWriter, r)

		// Log the request
		logger.Info("Request: %s %s %s %d %s",
			r.Method,
			r.RequestURI,
			r.Proto,
			wrappedWriter.Status(),
			time.Since(start))
	})

}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	// Default status code is 200 (OK),since the writer header is not called if the handler succeeds without the error
	return &ResponseWriter{
		ResponseWriter: w,
		status:         http.StatusOK,
	}
}

// WriteHeader captures the status code and delegates the call to the underlying ResponseWriter
func (rw *ResponseWriter) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}

// Status returns the status code of the response
func (rw *ResponseWriter) Status() int {
	return rw.status
}
