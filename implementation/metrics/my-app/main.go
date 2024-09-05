package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-faker/faker/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Device struct {
	ID       int    `json:"id"`
	Mac      string `json:"mac"`
	Firmware string `json:"firmware"`
}

// type PingResponse struct {
// 	Msg string `json:"msg"`
// }

type metrics struct {
	devices  prometheus.Gauge
	info     *prometheus.GaugeVec
	upgrades *prometheus.CounterVec
}

type registerDeviceHandler struct {
	metrics *metrics
}

func NewMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{
		devices: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "metrics",
			Name:      "connected_devices",
			Help:      "Number of currently connected devices",
		}),
		info: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "my_app",
			Name:      "info",
			Help:      "Information about app",
		},
			[]string{"version"}),
		upgrades: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "my_app",
			Name:      "device_upgrade_total",
			Help:      "Number of upgraded devices",
		}, []string{"type"}),
	}
	reg.MustRegister(m.devices, m.info, m.upgrades)
	return m
}

var version string
var dvs []Device

func createDevice(w http.ResponseWriter, r *http.Request, m *metrics) {
	var dv Device

	err := json.NewDecoder(r.Body).Decode(&dv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dvs = append(dvs, dv)
	m.devices.Set(float64(len(dvs)))

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Device created!"))
}

func getDevices(w http.ResponseWriter, _ *http.Request) {
	b, err := json.Marshal(dvs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (rdh registerDeviceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getDevices(w, r)
	case "POST":
		createDevice(w, r, rdh.metrics)
	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	reg := prometheus.NewRegistry()
	m := NewMetrics(reg)

	m.devices.Set(float64(len(dvs)))
	m.info.With(prometheus.Labels{"version": version}).Set(1)

	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})

	dMux := http.NewServeMux()
	rdh := registerDeviceHandler{metrics: m}
	dMux.HandleFunc("/devices", rdh.ServeHTTP)

	pMux := http.NewServeMux()
	pMux.Handle("/metrics", promHandler)

	go func() {
		log.Fatal(http.ListenAndServe(":8080", dMux))
	}()

	go func() {
		log.Fatal(http.ListenAndServe(":8081", pMux))
	}()

	select {}

}

// func ping(w http.ResponseWriter, req *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)

// 	b, err := json.Marshal(PingResponse{
// 		Msg: "Hello World",
// 	})
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.Write(b)

// }

func init() {
	version = "2.10.5"
	dvs = generateRandomDevices(100)
}

func generateRandomDevices(count int) []Device {
	var devices []Device

	for i := 0; i < count; i++ {
		v1, _ := faker.RandomInt(0, 9, 1)
		v2, _ := faker.RandomInt(0, 9, 1)
		v3, _ := faker.RandomInt(0, 9, 1)
		device := Device{
			ID:       i + 1, // Assign a unique ID starting from 1
			Mac:      faker.MacAddress(),
			Firmware: fmt.Sprintf("%d.%d.%d", v1[0], v2[0], v3[0]),
		}
		devices = append(devices, device)
	}

	return devices
}
