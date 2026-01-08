# HexaShop (Microservices + Architecture Hexagonale)

Projet exemple de microservices structurés en **architecture hexagonale (Ports & Adapters)**.  
Objectif : isoler le **domaine** (métier) des détails techniques (HTTP, DB, messaging), pour faciliter les tests, l’évolution et le remplacement des adapters.

---

## 🧭 Vision

- **Microservices** : chaque service est autonome (code, DB, déploiement).
- **Hexagonal** : le domaine ne dépend de rien.
- **DDD-friendly** : bounded contexts clairs (Customer, Order, Product, Payment…).

---

## 🧱 Microservices

- `customer-service` : gestion des clients et leurs adresses
- `product-service` : catalogue de produits et leur quantité en stock
- `order-service` : commandes, lignes de commande effectuées par les customers, 
- `payment-service` : service de paiements
- `gateway-proxy` : API Gateway point d’entrée dans l'univers des microservices de l'app

---

## 🗂️ Structure `customer-microservice`  en archi hexagonale
customer-microservice/
├── cmd/
│   └── api/
│       ├── main.go            # composition root (wiring)
│       ├── routes.go          # register routes (gin/nethttp)
│       └── container.go       # build dependencies (db, repos, usecases, handlers)
│
├── internal/
│   ├── domain/                # OBJETS MÉTIER (purs)
│   │   ├── customer.go        # objet métier Customer
│   │   ├── address.go         # objet métier Address
│   │   ├── valueobjects/      # Value Objects (validation métier)
│   │   │   ├── email.go
│   │   │   ├── phone.go
│   │   │   └── zipcode.go
│   │   └── errors.go          # erreurs métier
│   │
│   ├── application/           # USE CASES + PORTS
│   │   ├── ports/
│   │   │   ├── in/
│   │   │   │   ├── customer_uc.go
│   │   │   │   └── address_uc.go
│   │   │   └── out/
│   │   │       ├── customer_service.go
│   │   │       └── address_service.go
│   │   │
│   │   └── usecase/
│   │       ├── customer_service.go
│   │       └── address_service.go
│   │
│   ├── infrastructure/        # ADAPTERS (extérieur)
│   │   ├── web/
│   │   │   └── http/
│   │   │       ├── handlers/
│   │   │       │   ├── customer_handler.go
│   │   │       │   └── address_handler.go
│   │   │       ├── dtos/
│   │   │       │   ├── customer_request.go
│   │   │       │   ├── customer_response.go
│   │   │       │   ├── address_request.go
│   │   │       │   └── address_response.go
│   │   │       └── mappers/
│   │   │           ├── customer_mapper.go
│   │   │           └── address_mapper.go
│   │   │
│   │   └── persistence/
│   │       └── postgres/
│   │           ├── db.go
│   │           ├── models/
│   │           │   ├── customer_row.go
│   │           │   └── address_row.go
│   │           ├── mappers/
│   │           │   ├── customer_mapper.go
│   │           │   └── address_mapper.go
│   │           └── repositories/
│   │               ├── customer_repo.go
│   │               └── address_repo.go
│   │
│   ├── config/
│   │   ├── config.go
│   │   └── logger.go
│
├── migrations/
│   ├── 001_create_addresses.sql
│   └── 002_create_customers.sql
│
├── .gitignore
├── go.mod
└── README.md
