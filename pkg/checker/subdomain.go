package checker

import (
	"net/http"
	"strings"

	"github.com/mugund10/simpleserver/pkg/readers"
)

// a middleware which checks for subdomain
func CheckSubdomain(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Host == "mugund10.dev" {
			next.ServeHTTP(w, r)
		} else {
			// checks domain too
			sd := readers.GetServerS()
			checks := sd[0].Domain
			if strings.Contains(r.Host, checks) {
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
		}
	})
}

// finds whether the requested host and config host are same
func finder(match string) bool {
	Proxies := readers.Getproxies()
	for i := 0; i < len(Proxies); i++ {
		if Proxies[i].Subdomain == match {
			return true
		}
	}
	return false
}
