# HexaShop : Microservices + Archi Hexagonale

Projet exemple de microservices structurés en **architecture hexagonale (Ports & Adapters)**.  
Objectif : isoler le **domaine** (métier) des détails techniques (HTTP, DB, messaging), pour faciliter les tests, l’évolution et le remplacement des adapters.

## 🧭 Vision

- **Microservices** : chaque service est autonome (code, DB, déploiement).
- **Hexagonal** : le domaine ne dépend de rien.
- **DDD-friendly** : bounded contexts clairs (Customer, Order, Product, Payment…).

## 🧱 Microservices

- `customer-service` : gestion des clients et leurs adresses
- `product-service` : catalogue de produits et leur quantité en stock
- `order-service` : commandes, lignes de commande effectuées par les customers, 
- `payment-service` : service de paiements
- `gateway-proxy` : API Gateway point d’entrée dans l'univers des microservices de l'app

## 🗂️ Structure `customer-microservice`  en archi hexagonale
```
customer-microservice/
├── cmd/
│   └── api/
│       ├── main.go                                         # composition root (wiring)
│       ├── routes.go                                       # register routes (gin/nethttp)
│       └── container.go                                    # build dependencies (db, repos, usecases, handlers)
│
├── internal/
│   ├── domain/                                             # 1️⃣ OBJETS MÉTIER (purs)
│   │   ├── business_customer.go                            # objet métier Customer
│   │   ├── business_address.go                             # objet métier Address
│   │   ├── validator/          
│   │   │   └── fieds_checker.go                            # validation des champs du métier
│   │   └── errors.go                                       # erreurs métier
│   │   │
│   ├── application/                                        # 2️⃣ USE CASES + PORTS
│   │   ├── ports/
│   │   │   ├── in/
│   │   │   │   ├── customer_in_port.go                     # InCustomerService port d'entrée exposé à l'extreieur
|   |   |   |   └── address_in_port.go                      # InAddressService port d'entrée exposé à l'extreieur
│   │   │   ├── out/
|   |   |   │   ├── customer_out_port.go                    # OutCustomerService utilisé pour envoyer à l'exterieur
|   |   |   |   └── address_out_port.go                     # OutAddressService utilisé pour envoyer à l'exterieur
│   │   │   └── usecase/                                    # ✅ usecase implemente les input ports
│   │   │       ├── customer_uc.go
│   │   │       └── address_uc.go
│   │   │
│   ├── infrastructure/                                     # 3️⃣ ADAPTERS (extérieur)
│   │   ├── web/
│   │   │   └── http/
│   │   │       ├── handlers/
|   |   |       |   ├── contract/                           # hanlder avec gin-gonic
│   │   │       │   |    ├── customer_handler.go            # interface CustomerHandlerService 
│   │   │       │   |    └── address_handler.go             # interface AddressHandlerService 
|   |   |       |   ├── impl/ 
|   |   |       |        ├── customer_handler_impl.go       # impl CustomerHandlerService 
│   │   │       │        └── address_handler_impl.go        # implAddressHandlerService        
|   |   |       ├── routes/
|   |   |       |   └── route_register.go                   # engeristrement des routes avec gin.Engine                                 
│   │   │       ├── dtos/                                   # ✅  les user dtos                             
│   │   │       │   ├── customer_request.go
│   │   │       │   ├── customer_response.go
│   │   │       │   ├── address_request.go
│   │   │       │   └── address_response.go
│   │   │       └── mappers/                                # ✅ mappers de transformation
│   │   │           ├── customer_mapper.go
│   │   │           └── address_mapper.go
│   │   │
│   │   ├──  persistence/                                   # ✅ save dans la db
│   │   |       └── postgres/                       
│   │   |           ├── db.go                               # db *sql.DB par exemple
│   │   |           ├── models/
│   │   |           │   ├── customer_table.go               # model de données de la table customers
│   │   |           │   └── address_table.go                # model de données de la table addresses
│   │   |           ├── mappers/
│   │   |           │   ├── customer_mapper.go
│   │   |           │   └── address_mapper.go
│   │   |           └── repositories/                       # ✅ implementation des outputs ports
│   │   |               ├── customer_out_port_impl.go       # impl du customer output port
|   |   |               └── address_out_port_impl.go        # impl de address output port
│   │   |
│   ├── config/                                             # 4️⃣ la config des env vars
│   │   └── config.go
│   │
├── migrations/
│   ├── 001_create_addresses.sql
│   └── 002_create_customers.sql
│
├── .gitignore
├── go.mod
└── README.md
```
