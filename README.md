# iot-air-monitoring

Dashboard sederhana untuk memantau kualitas udara dari IoT device.

Alur: IoT device → `POST /api/sensor-data` → Go → MySQL → dashboard `/sensor`

## Cara menjalankan

1. Buat database dan tabel (sekali saja):
   ```bash
   mysql -u root -p < schema.sql
   ```
2. Sesuaikan user dan password MySQL di `cmd/database.go`.
3. Jalankan server:
   ```bash
   go run ./cmd
   ```
4. Buka http://localhost:8080/sensor

## Contoh kirim data

```bash
curl -X POST localhost:8080/api/sensor-data -d '{"device_id":101,"lat":-6.9147,"long":107.6098,"temperature":28.5,"pressure":1010.2}'
```
