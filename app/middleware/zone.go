package middleware

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"

	"dnsgo/app/cache"
)

var Zone = &zoneMiddle{}

type zoneMiddle struct{}

func (m *zoneMiddle) GetZone(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		zoneID := mux.Vars(r)["zone_id"]
		zone, err := cache.Zone.Get(r.Context(), zoneID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, "zone", zone)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
