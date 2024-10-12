# Test Backend Anekapay

Proyek ini adalah backend sederhana untuk aplikasi Anekapay yang dibangun menggunakan Go dengan framework Echo. Proyek ini mencakup pengaturan, inisialisasi, dan migrasi database.

## Prerequisites

Sebelum memulai, pastikan Anda telah menginstal:

- [Go](https://golang.org/doc/install) (versi 1.18 atau lebih baru)
- [Docker](https://docs.docker.com/get-docker/) (opsional, jika Anda ingin menjalankan menggunakan Docker)
- [MySQL](https://www.mysql.com/downloads/) atau MySQL untuk database

## Clone Proyek
Clone repositori ini ke dalam direktori lokal Anda:
```bash
git clone https://github.com/username/test-anekapay-backend.git
```

# Masuk ke direktori proyek
```bash
cd test-anekapay-backend
```
# Inisialisasi modul Go
```bash
go mod init test-anekapay-backend
```
# Instal dependensi Echo
go get github.com/labstack/echo/v4

# File Database dan Migrasi ada pada Folder ../config/migration/ .sql

# API Endpoint
## Silakan lihat dokumentasi API untuk daftar endpoint yang tersedia dan cara penggunaannya. Ada pada link berikut: 
(https://documenter.getpostman.com/view/30196710/2sAXxS8BHa).
