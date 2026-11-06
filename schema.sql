-- Jalankan sekali saja di MySQL untuk membuat database dan tabel
CREATE DATABASE IF NOT EXISTS iot_monitoring;
USE iot_monitoring;

-- Daftar device. Harus dibuat lebih dulu karena sensor_readings menunjuk ke tabel ini
CREATE TABLE IF NOT EXISTS devices (
    id           INT PRIMARY KEY,
    name         VARCHAR(100),
    city         VARCHAR(100),
    installed_at DATE
);

CREATE TABLE IF NOT EXISTS sensor_readings (
    id          INT AUTO_INCREMENT PRIMARY KEY,
    device_id   INT,
    latitude    DOUBLE,
    longitude   DOUBLE,
    temperature DOUBLE,
    pressure    DOUBLE,
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,

    -- Foreign key: device_id wajib ada di tabel devices
    CONSTRAINT fk_sensor_readings_device FOREIGN KEY (device_id) REFERENCES devices(id)
);

-- Data awal device, supaya device bisa langsung mengirim data
INSERT IGNORE INTO devices (id, name, city, installed_at) VALUES
(101, 'Sensor Gedung Sate',   'Bandung',  '2026-10-01'),
(102, 'Sensor Monas',         'Jakarta',  '2026-10-15'),
(103, 'Sensor Tugu Pahlawan', 'Surabaya', '2026-10-20'),
(104, 'Sensor Renon',         'Denpasar', '2026-11-05');
