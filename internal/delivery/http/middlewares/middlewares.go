package middlewares

import (
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

			// TODO Header

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
