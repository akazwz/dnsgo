package api

import (
	"encoding/json"
	"net/http"
	"time"

	"dnsgo/app/model"
	"dnsgo/app/service"
	"dnsgo/utils"
)

var Zone = &zoneApi{}

type zoneApi struct{}

func (a *zoneApi) List(w http.ResponseWriter, r *http.Request) {
	zones, err := service.Zone.Find(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(zones)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *zoneApi) Get(w http.ResponseWriter, r *http.Request) {
	zone := r.Context().Value("zone").(*model.Zone)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(zone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *zoneApi) Create(w http.ResponseWriter, r *http.Request) {
	zone := &model.Zone{}
	err := json.NewDecoder(r.Body).Decode(zone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := utils.ID.Generate(21)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	zone.ID = id
	now := time.Now().UTC()
	zone.CreatedAt = now
	zone.UpdatedAt = now

	err = service.Zone.Create(r.Context(), zone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(zone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *zoneApi) Update(w http.ResponseWriter, r *http.Request) {
	zone := r.Context().Value("zone").(*model.Zone)
	id := zone.ID
	err := json.NewDecoder(r.Body).Decode(zone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	zone.ID = id
	zone.UpdatedAt = time.Now().UTC()

	err = service.Zone.Update(r.Context(), zone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(zone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *zoneApi) Delete(w http.ResponseWriter, r *http.Request) {
	zone := r.Context().Value("zone").(*model.Zone)
	id := zone.ID
	err := service.Zone.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
