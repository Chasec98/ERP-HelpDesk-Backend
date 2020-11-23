package pagination

import (
	"context"
	"net/http"
	"strconv"
)

func PaginationCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			http.Error(w, "provide valid limit", http.StatusBadRequest)
			return
		}
		offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err != nil {
			http.Error(w, "provide valid offset", http.StatusBadRequest)
			return
		}

		var paginationParams = PaginationContext{
			Limit:  limit,
			Offset: offset,
		}

		ctx := context.WithValue(r.Context(), PaginationCtxKey, paginationParams)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
