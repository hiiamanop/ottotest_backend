🚀 Overview

OTTO DIGITAL test Backend ini adalah layanan RESTful API berbasis Golang (clean architecture) untuk proses redemption voucher dengan sistem point, mendukung multi-brand, multi-voucher, customer, serta atomic transaction. Proyek ini dibangun dengan Gin, GORM, PostgreSQL, Docker, dan sepenuhnya terdokumentasi Swagger

🧩 Tech Stack

- Go (Golang)
- Gin (HTTP framework)
- GORM (ORM)
- PostgreSQL
- Docker
- Swagger (API Docs, via swaggo/swag)
- GoMock & Testify (unit testing)

📖 Dokumentasi API (Swagger)

Swagger UI bisa diakses di:
http://localhost:8080/swagger/index.html

🔗 Contoh Endpoint

1. Customer

   - POST /customer

   {
   "name": "Budi",
   "email": "budi@mail.com",
   "point_balance": 200000
   }

   - GET /customer?id=1

2. Brand
   - POST /brand
     {
     "name": "Indomaret",
     "description": "Minimarket"
     }
3. Voucher

   - POST /voucher

   {
   "brand_id": 1,
   "name": "Voucher 50K",
   "description": "Voucher belanja",
   "point_cost": 50000
   }

   - GET /voucher?id=1

   - GET /voucher/brand?id=1

4. Redemption

   - POST /transaction/redemption

   {
   "customer_id": 1,
   "vouchers":
   [
   { "voucher_id": 1, "quantity": 2 }
   ]
   }

   - GET /transaction/redemption?transactionId=1

🧪 Unit Testing

Menggunakan GoMock dan Testify untuk isolasi business logic

🎯 Fitur Utama

1. Multi-brand & multi-voucher
2. Redeem voucher dengan validasi point customer
3. Atomic transaction (unit tested)
4. Clean architecture (maintainable, scalable)
5. Dokumentasi API otomatis (Swagger)
6. Dockerized + migration otomatis

📝 Catatan

Semua endpoint dan struktur data lengkap dapat dilihat di Swagger.
