package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"iot-air-monitoring/web"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Satu baris data yang dikirim oleh IoT device
type Reading struct {
	DeviceID    int     `json:"device_id"`
	Timestamp   string  `json:"timestamp"`
	Lat         float64 `json:"lat"`
	Long        float64 `json:"long"`
	Temperature float64 `json:"temperature"`
	Pressure    float64 `json:"pressure"`
}

func main() {
	ConnectDB()

	router := chi.NewRouter()

	router.Get("/", Home)
	router.Get("/home", Home)
	router.Get("/sensor", Sensor)
	router.Post("/api/sensor-data", SaveSensorData)

	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", router)
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(
		web.Files,
		"template/base.html",
		"home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Sensor(w http.ResponseWriter, r *http.Request) {
	// Data dummy dulu, nanti diganti dengan SELECT dari MySQL di bawah
	readings := []Reading{
		{101, "2026-10-26 20:05:00", -6.9147, 107.6098, 31.2, 1009.8},
		{101, "2026-10-26 20:04:00", -6.9147, 107.6098, 30.4, 1010.1},
		{101, "2026-10-26 20:03:00", -6.9147, 107.6098, 29.6, 1010.3},
		{101, "2026-10-26 20:02:00", -6.9147, 107.6098, 28.9, 1010.6},
		{101, "2026-10-26 20:01:00", -6.9147, 107.6098, 28.1, 1010.9},
		{101, "2026-10-26 20:00:00", -6.9147, 107.6098, 27.5, 1011.2},
	}

	// Ambil 50 data terbaru dari MySQL (hapus data dummy di atas kalau ini dipakai)
	// rows, err := DB.Query(
	// 	"SELECT device_id, created_at, latitude, longitude, temperature, pressure FROM sensor_readings ORDER BY id DESC LIMIT 50")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// defer rows.Close()
	//
	// readings := []Reading{}
	// for rows.Next() {
	// 	var reading Reading
	// 	rows.Scan(&reading.DeviceID, &reading.Timestamp, &reading.Lat, &reading.Long, &reading.Temperature, &reading.Pressure)
	// 	readings = append(readings, reading)
	// }

	// Data terbaru ada di baris pertama (kalau tabel masih kosong, pakai nilai 0)
	latest := Reading{}
	if len(readings) > 0 {
		latest = readings[0]
	}

	tmpl, err := template.ParseFS(
		web.Files,
		"template/base.html",
		"sensor.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Kirim data terbaru dan semua data historis ke HTML
	data := map[string]any{
		"Latest":   latest,
		"Readings": readings,
	}

	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Dipanggil oleh IoT device: terima JSON, lalu simpan ke MySQL
func SaveSensorData(w http.ResponseWriter, r *http.Request) {
	var reading Reading
	json.NewDecoder(r.Body).Decode(&reading)

	_, err := DB.Exec(
		"INSERT INTO sensor_readings (device_id, latitude, longitude, temperature, pressure) VALUES (?, ?, ?, ?, ?)",
		reading.DeviceID, reading.Lat, reading.Long, reading.Temperature, reading.Pressure)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("OK"))
}
