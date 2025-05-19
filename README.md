# SE4458 Go API Gateway

Bu proje, Laravel tabanlı iki mikroservisi (App A - GSM Operator API, App B - ChatBot API) yöneten bir **API Gateway** servisidir. Go ile yazılmıştır ve tüm istemci istekleri bu gateway üzerinden yönlendirilir.

## 🛠 Özellikler

- Go (net/http + chi router)
- Reverse Proxy mantığı
- `.env` ile yapılandırılabilir
- App A ve App B isteklerini ayrı yönlendirir

---

## 📁 Proje Yapısı

```
go-gateway/
├── main.go             // Uygulama başlangıcı
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

# Laravel App A (GSM Operator API)
GSM_API_URL=

# Laravel App B (ChatBot API)
CHAT_API_URL=
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

### GSM API (App A):

```http
POST http://localhost:8080/gsm/api/v1/subscriber
```

Bu istek doğrudan App A’ya yönlendirilir:

```
-> http://xellpay.test/api/v1/subscriber
```

### ChatBot API (App B):

```http
POST http://localhost:8080/chat/api/v1/chat
```

Bu istek de App B’ye yönlendirilir:

```
-> http://se4458-chatapp.test/api/v1/chat
```

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
