# Go Image Uploader

Aplikasi RESTful API untuk upload dan mengelola gambar. Dibuat dengan Go Fiber dan PostgreSQL.

## Clone Repository

```bash
# Clone repository
git clone https://github.com/MapIHS/go-image-uploader

# Pindah ke direktori proyek
cd go-image-uploader
```

## Instalasi Cepat

```bash
# Download dependencies
go mod download

# run projek
go run ./cmd/server

# Build aplikasi
go build -o app ./cmd/server

# Jalankan aplikasi
./app
```

## Konfigurasi Minimal

Buat file `.env` di root proyek:

```
PORT=8080
IMAGE_SIZE_LIMIT=10485760
UPLOAD_PATH=/path/to/images
POSTGRES_CONN=postgresql://username:password@localhost:5432/imagedb?sslmode=disable
```

## API Endpoints

### Upload dan Akses Gambar

| Metode | Endpoint | Deskripsi | Parameter |
|--------|----------|-----------|-----------|
| GET | /api/images | Mendapatkan semua gambar | - |
| POST | /api/upload | Upload gambar baru | form-data: `image` |

### Status & Pemantauan

| Metode | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | /api/disk-usage | Mendapatkan informasi penggunaan disk |

## Contoh Request

### Upload Gambar
```bash
curl -X POST http://localhost:8080/api/upload \
  -H "Content-Type: multipart/form-data" \
  -F "image=@/path/to/your/image.jpg"
```

### Mendapatkan Gambar
```bash
curl -X GET http://localhost:8080/api/images
```

### Penggunaan Disk
```bash
curl -X GET http://localhost:8080/api/disk-usage
```