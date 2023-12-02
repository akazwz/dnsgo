package api

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"dnsgo/app/model"
	"dnsgo/app/service"
	"dnsgo/utils"
)

var DNSRecord = &dnsRecordApi{}

type dnsRecordApi struct{}

func (a *dnsRecordApi) List(w http.ResponseWriter, r *http.Request) {
	zone := r.Context().Value("zone").(*model.Zone)
	records, err := service.DNSRecord.GetZoneDNSRecords(r.Context(), zone.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(records)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *dnsRecordApi) Get(w http.ResponseWriter, r *http.Request) {
	record := r.Context().Value("record").(*model.DNSRecord)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *dnsRecordApi) Create(w http.ResponseWriter, r *http.Request) {
	zone := r.Context().Value("zone").(*model.Zone)
	record := &model.DNSRecord{}
	err := json.NewDecoder(r.Body).Decode(record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, err := utils.ID.Generate(21)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	record.ID = id
	record.ZoneID = zone.ID
	record.ZoneName = zone.Name
	name := record.Name
	if !strings.HasSuffix(name, zone.Name) {
	}
	now := time.Now().UTC()
	record.CreatedAt = now
	record.UpdatedAt = now
	err = service.DNSRecord.CreateDNSRecord(r.Context(), record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (a *dnsRecordApi) Update(w http.ResponseWriter, r *http.Request) {
	record := r.Context().Value("record").(*model.DNSRecord)
	id := record.ID
	err := json.NewDecoder(r.Body).Decode(record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	record.ID = id
	record.UpdatedAt = time.Now().UTC()
	err = service.DNSRecord.UpdateDNSRecord(r.Context(), record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (a *dnsRecordApi) Delete(w http.ResponseWriter, r *http.Request) {
	record := r.Context().Value("record").(*model.DNSRecord)
	id := record.ID
	err := service.DNSRecord.DeleteDNSRecord(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (a *dnsRecordApi) DeleteByZoneId(w http.ResponseWriter, r *http.Request) {
	zone := r.Context().Value("zone").(*model.Zone)
	err := service.DNSRecord.DeleteDNSRecordByZoneId(r.Context(), zone.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
