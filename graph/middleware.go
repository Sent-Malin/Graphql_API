package graph

import (
	"context"
	"net/http"

	"github.com/go-pg/pg/v10"
)

//Middleware возвращает handler, записывающий
//токен авторизации в контекст
func Middleware(db *pg.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c := r.Header.Get("Authorization")

			if c == "" {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), "authorization", c)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

//ForContext находит jwt-токен в контексте
func ForContext(ctx context.Context) string {
	raw, _ := ctx.Value("authorization").(string)
	return raw
}
