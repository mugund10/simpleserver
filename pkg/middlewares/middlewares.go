package middlewares

import "net/http"


type Mware func(http.Handler) http.Handler


func MakeStack(xs ...Mware) Mware {
	return func(next http.Handler) http.Handler {
		for _, x := range xs {
			next = x(next)
		}
		return next
	}
}
