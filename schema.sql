-- Jalankan sekali saja di MySQL untuk membuat database dan tabel
CREATE DATABASE IF NOT EXISTS iot_monitoring;
USE iot_monitoring;

CREATE TABLE IF NOT EXISTS sensor_readings (
    id          INT AUTO_INCREMENT PRIMARY KEY,
    device_id   INT,
    latitude    DOUBLE,
    longitude   DOUBLE,
    temperature DOUBLE,
    pressure    DOUBLE,
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP
);
