package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/eneridangelis/device-rest/internal/model"
	"github.com/eneridangelis/device-rest/internal/usecase"
	"github.com/gorilla/mux"
)

func NewRouter(deviceHandler *DeviceHandler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/devices", deviceHandler.AddDevice).Methods("POST")
	router.HandleFunc("/devices/search", deviceHandler.SearchDevices).Methods("GET")
	router.HandleFunc("/devices/{id}", deviceHandler.GetDeviceByID).Methods("GET")
	router.HandleFunc("/devices", deviceHandler.ListDevices).Methods("GET")
	router.HandleFunc("/devices", deviceHandler.UpdateDevice).Methods("PATCH")
	router.HandleFunc("/devices/{id}", deviceHandler.DeleteDevice).Methods("DELETE")
	return router
}

type DeviceHandler struct {
	usecase *usecase.DeviceUsecase
}

func NewDeviceHandler(usecase *usecase.DeviceUsecase) *DeviceHandler {
	return &DeviceHandler{usecase: usecase}
}

func (h *DeviceHandler) AddDevice(w http.ResponseWriter, r *http.Request) {
	var device *model.Device
	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.usecase.AddDevice(device); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *DeviceHandler) GetDeviceByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid device ID", http.StatusBadRequest)
		return
	}

	device, err := h.usecase.GetDeviceByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(device)
}

func (h *DeviceHandler) ListDevices(w http.ResponseWriter, r *http.Request) {
	devices, err := h.usecase.ListAllDevices()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(devices)
}

func (h *DeviceHandler) UpdateDevice(w http.ResponseWriter, r *http.Request) {
	var device *model.Device
	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.usecase.UpdateDevice(device); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *DeviceHandler) DeleteDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid device ID", http.StatusBadRequest)
		return
	}

	if err := h.usecase.DeleteDevice(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *DeviceHandler) SearchDevices(w http.ResponseWriter, r *http.Request) {
	brand := r.URL.Query().Get("brand")
	devices, err := h.usecase.SearchDeviceByBrand(brand)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(devices)
}
