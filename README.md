# SE4458 Go API Gateway

Bu proje, Laravel tabanlı iki mikroservisi (App A - Auth API, App B - Job Posting API) yöneten bir **API Gateway** servisidir. Go ile yazılmıştır ve tüm istemci istekleri bu gateway üzerinden yönlendirilir.

## 🛠 Özellikler

- Go (net/http + chi router)
- Reverse Proxy mantığı
- CORS desteği (localhost:5173 için)
- `.env` ile yapılandırılabilir
- Auth ve Job Posting isteklerini ayrı yönlendirir

---

## 📁 Proje Yapısı

```
go-gateway/
├── main.go             // Uygulama başlangıcı ve CORS middleware
├── handlers/
│   └── proxy.go        // Reverse proxy logic
├── .env.example        // Örnek Ortam değişkenleri
├── go.mod              // Go modül yapılandırması
└── README.md
```

---

## ⚙️ .env Dosyası

```env
PORT=8080

AUTH_API_URL=

JOB_POSTING_API_URL=
```

---

## 🚀 Kurulum ve Çalıştırma

### 1. Gerekli Go paketlerini indir:

```bash
go mod tidy
```

### 2. Projeyi çalıştır:

```bash
go run main.go
```

Gateway şu adreste çalışır:  
`http://localhost:8080`

---

## 🧪 Kullanım

### Auth API (App A):

```http
POST http://localhost:8080/auth/api/v1/login
```

Bu istek doğrudan Auth API'ye yönlendirilir.

### Job Posting API (App B):

```http
GET http://localhost:8080/jobs/api/v1/jobs?page=1&limit=10
POST http://localhost:8080/jobs/api/v1/jobs
```

Bu istekler Job Posting API'ye yönlendirilir.

### Job Search API:

```http
GET http://localhost:8080/job-search/api/v1/jobs/search?q=keyword
```

Bu istek de Job Posting API'nin search endpoint'ine yönlendirilir.

---

## 🌐 CORS Desteği

Gateway, frontend uygulamalarından gelen istekleri desteklemek için CORS middleware'i içerir:

- **İzin Verilen Origin:** `http://localhost:5173`
- **İzin Verilen Metodlar:** GET, POST, PUT, DELETE, OPTIONS
- **İzin Verilen Headers:** Content-Type, Authorization

---

## 🧩 Supervisor ile Servis Olarak Çalıştırmak

```ini
[program:go-gateway]
directory=/var/www/se4458-go-gateway
command=/usr/bin/env go run main.go
autostart=true
autorestart=true
startsecs=5
user=root
environment=GOPATH="/root/go",GOMODCACHE="/root/go/pkg/mod",GOCACHE="/tmp/go-cache",PATH="/usr/local/go/bin:/usr/bin:/bin"
stdout_logfile=/var/log/go-gateway.log
stderr_logfile=/var/log/go-gateway-error.log
```

---

## 📦 Bağımlılıklar

- Go 1.20+
- [go-chi/chi](https://github.com/go-chi/chi)
- [joho/godotenv](https://github.com/joho/godotenv)
