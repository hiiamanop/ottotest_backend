# 🚀 Voucher Redemption Service (Backend Test)

Backend API service for voucher redemption, built with Go (Golang), Clean Architecture, PostgreSQL, Docker, and full unit testing.  
Submission for Backend Developer Test.

---

## ✨ Features

- Clean/Hexagonal Architecture
- Voucher-Brand, Customer, Transaction entities
- Redemption transaction logic with point validation
- Input validation & error handling
- Database migration
- Unit tests with gomock
- Docker & docker-compose ready
- API documentation with Swagger (OpenAPI)
- Environment config via .env

---

## 🚀 Quick Start

**1. Clone Repository & Setup .env**

```bash
git clone https://github.com/your-username/ottotest_backend.git
cd ottotest_backend
cp .env.example .env  # Edit if needed (DB credentials, etc)
```

**2. Build & Run all services (App, DB, Migration)**

```bash
docker-compose up --build
```

App: http://localhost:8080  
Swagger: http://localhost:8080/swagger/index.html  
DB: localhost:5432 (ottotest_db)

**3. Run Database Migration (if not auto-run)**

```bash
docker-compose run migrate
```

**4. Run Unit Tests (in local dev)**

```bash
go test ./internal/...
```

## 🧪 API Endpoints

- POST /brand - Create brand
- POST /voucher - Create voucher
- GET /voucher?id={voucher_id} - Get voucher by ID
- GET /voucher/brand?id={brand_id} - Get all vouchers by brand
- POST /transaction/redemption - Create redemption transaction (multi-voucher)
- GET /transaction/redemption?transactionId={id} - Get transaction detail
- POST /customer - Create customer
- GET /customer?id={id} - Get customer by ID

See Swagger UI for full documentation & schema.

## 🗂️ Project Structure

```
voucher-redemption-service/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── domain/
│   │   ├── entity/
│   │   └── repository/
│   ├── usecase/
│   │   └── service/
│   ├── infrastructure/
│   │   ├── persistence/
│   │   └── http/
│   └── interface/
│       ├── handler/
│       └── dto/
├── pkg/
│   ├── config/
│   └── util/
├── migrations/
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

## ⚙️ Environment Variables

Set in .env (see .env.example):

- DATABASE_URL
- APP_PORT
- etc

## 🧪 Testing

Unit Tests (with gomock):

```bash
go test ./internal/...
```

Run inside Docker:

```bash
docker-compose exec app go test ./internal/...
```

## 📝 Notes

- Default PostgreSQL credentials can be changed in docker-compose.yml and .env.
- Migrations auto-run on docker-compose up (can be run manually if needed).
- See migrations/ for SQL schema.

## 👤 Author

Ahmad Naufal Muzakki

## 📄 License

MIT
EOF
