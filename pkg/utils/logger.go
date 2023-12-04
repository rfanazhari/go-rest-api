package utils

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	// You can set the desired log level (e.g., Info, Debug, Error)
	logger.SetLevel(logrus.InfoLevel)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log information for each incoming request
		logger.WithFields(logrus.Fields{
			"method": r.Method,
			"path":   r.URL.Path,
			"ip":     r.RemoteAddr,
		}).Info("incoming")

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
