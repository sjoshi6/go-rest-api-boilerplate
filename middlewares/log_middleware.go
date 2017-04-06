package middlewares

import (
	"net/http"
	"time"

	"github.com/sjoshi6/go-rest-api-boilerplate/logger"
	"github.com/urfave/negroni"
)

// LoggingMiddleWare is used to depict a request logger
type LoggingMiddleWare struct {
}

// ServeHTTP implements the negroni middleware function
func (lm *LoggingMiddleWare) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	t1 := time.Now()
	next(rw, r)
	t2 := time.Now()

	res := rw.(negroni.ResponseWriter)
	log := logger.GetLogger()
	log.Infof("%s %s %s %d %s %s %s %v", r.Host, r.Method, r.URL.Path, res.Status(), r.Proto, r.RemoteAddr, r.Header.Get("X-Forwarded-For"), t2.Sub(t1))
}
