package checker

import (
	"net/http"
	"strings"

	"github.com/mugund10/simpleserver/pkg/reverse"
)

// a middleware which checks for subdomain
func CheckSubdomain(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// checks domain too
		if strings.Contains(r.Host, "mugund10.top:") {
			sdom := strings.Split(r.Host, ".")
			ok := finder(sdom[0])
			if ok {
				next.ServeHTTP(w, r)
			} else {
				http.NotFound(w, r)
			}
		} else {
			http.NotFound(w, r)
		}
	})
}

// finds whether the requested host and config host are same
func finder(match string) bool {
	Proxies := reverse.Getproxies()
	for i := 0; i < len(Proxies); i++ {
		if Proxies[i].Subdomain == match {
			return true
		}
	}
	return false
}
