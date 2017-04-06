package middlewares

import (
	"net/http"

	"github.com/sjoshi6/go-rest-api-boilerplate/logger"
)

var validContentTypes = []string{"application/json"}

// ContentCheckMiddleware is used to check if content-Type of the POST Request is an acceptable type.
type ContentCheckMiddleware struct {
}

// CheckContentType is used to check the real contents of Content-type in http header
func CheckContentType(r *http.Request) bool {

	if r.Method == http.MethodPost {

		incomingContentType := r.Header.Get("Content-type")
		for _, contentType := range validContentTypes {
			if contentType == incomingContentType {
				return true
			}
		}
		return false
	}
	return true
}

// ServeHTTP implements the negroni middleware function
func (c *ContentCheckMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	if CheckContentType(r) == false {

		log := logger.GetLogger()
		log.Warnf("Received an invalid content-type in header. API only accepts %s", validContentTypes)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	next(rw, r)
}
