package middlewares

import (
	"fmt"
	"net/http"

	"github.com/andrsj/feedback-service/internal/infrastructure/cache"
)

func CacheMiddleware(cache cache.Cache) (func(next http.Handler) http.Handler) {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				next.ServeHTTP(w, r)

				return
			}

			cacheKey := r.URL.String()
			if val, ok := cache.Get(cacheKey); ok {
				w.Write(val) //nolint

				return
			}

			rw := NewResponseWriter(w, http.StatusProcessing)
			next.ServeHTTP(rw, r)

			if rw.Status() == http.StatusOK {
				cache.Set(cacheKey, rw.Body.Bytes())
			}
		})
	}
}

// TODO Remove this example of closures in middleware
func Mid (i int) (func(next http.Handler) http.Handler) {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("I-Custom-Header", fmt.Sprint(i))
			next.ServeHTTP(w, r)
		})
	}
}