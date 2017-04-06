package middlewares

import (
	"net/http"

	"github.com/sjoshi6/go-rest-api-boilerplate/logger"
)

var validContentTypes = []string{"application/json"}

type ContentCheckMiddleware struct {
}

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

func (c *ContentCheckMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	if CheckContentType(r) == false {

		log := logger.GetLogger()
		log.Warnf("Received an invalid content-type in header. API only accepts %s", validContentTypes)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	next(rw, r)
}
