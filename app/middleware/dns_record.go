package middleware

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"

	"dnsgo/app/cache"
)

var DNSRecord = &dnsRecordMiddle{}

type dnsRecordMiddle struct{}

func (m *dnsRecordMiddle) GetDNSRecord(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["record_id"]
		record, err := cache.DNSRecord.Get(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, "record", record)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
