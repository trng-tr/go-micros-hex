# HexaShop : Microservices + Archi Hexagonale

Projet exemple de microservices structurés en **architecture hexagonale (Ports & Adapters)**.  
Objectif : isoler le **domaine** (métier) des détails techniques (HTTP, DB, messaging), pour faciliter les tests, l’évolutivité et le remplacement des adapters, etc.

## 🧭 Vision

- **Microservices** : chaque service est autonome (code, DB, déploiement).
- **Hexagonal** : le domaine ne dépend de rien.
- **DDD-friendly** : bounded contexts clairs (Customer, Order, Product, Payment…).

## 🧱 Microservices

- `customer-microservice` : gestion des clients et leurs adresses
- `product-microservice` : catalogue de produits et leur quantité en stock
- `order-microservice` : commandes, lignes de commande effectuées par les customers, 
- `payment-microservice` : microservice de paiements
- `gateway-proxy` : API Gateway point d’entrée dans l'univers des microservices de l'app

## 🗂️ Structure `customer-microservice`  en archi hexagonale
```
customer-microservice/
├── cmd/
│   └── api/
│       ├── main.go                                         # composition root (wiring)
│
├── internal/
│   ├── domain/                                             # 1️⃣ OBJETS MÉTIER (purs)
│   │   ├── business_customer.go                            # BusinessCustomer objet métier Customer
│   │   ├── business_address.go                             # BusinessAddress objet métier Address
│   │   ├── validator/          
│   │   │   ├── fieds_checker.go                            # validation des champs du métier
│   │   |   └── errors.go                                       # erreurs métier
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
│   │   │       ├── customer_usecase.go
│   │   │       └── address_usecase.go
│   │   │
│   ├── infrastructure/                                     # 3️⃣ ADAPTERS (extérieur)         
│   │   ├── in/
│   │   │   └── web/
│   │   │       ├── handlers/
|   |   |       |   ├── customer_handler_impl.go            # impl CustomerHandlerService 
│   │   │       │   └── address_handler_impl.go             # implAddressHandlerService        
|   |   |       ├── routes/                                 # register routes 
|   |   |       |   ├── customer_handler.go                 # interface CustomerHandlerService: gin-gonic
|   |   |       |   ├── address_handler.go                  # implAddressHandlerService : gin-gonic   
|   |   |       |   └── route_register.go                   # engeristrement des routes: gin                        
│   │   │       ├── dtos/                                   # ✅  les user dtos                             
│   │   │       │   ├── customer_request.go
│   │   │       │   ├── customer_response.go
│   │   │       │   ├── address_request.go
│   │   │       │   └── address_response.go
│   │   │       └── mappers/                                # ✅ mappers de transformation
│   │   │           ├── customer_mapper.go
│   │   │           └── address_mapper.go
│   │   │
│   │   ├── out/                                           # ✅ save dans la db
│   │   |       └── services/                       
│   │   |           ├── db.go                               # db *sql.DB par exemple
│   │   |           ├── models/
│   │   |           │   ├── customer_model.go               # model de données pour la table customers
│   │   |           │   └── address_model.go                # model de données pour la table addresses
│   │   |           ├── mappers/
│   │   |           │   ├── customer_mapper.go
│   │   |           │   └── address_mapper.go
|   |   |           ├── repositories                        # ✅ la couche de données (db)
|   |   |           |   ├── customer_address_repos_impl.go  # implementation des customer et address repos 
│   │   |           └── services/                           # ✅ implementation des outputs ports
|   |   |               ├── generic_repos.go                # repo generic de centralisation des méthodes
|   |   |               ├── customer_address_repos.go       # real repo extends generic repo 
│   │   |               ├── customer_out_port_impl.go       # OutCustomerServiceImpl impl du customer output port
|   |   |               └── address_out_port_impl.go        # OutAddressServiceImpl impl de address output port
│   │   |
│   |   ├── config/                                          # 4️⃣ la config des env vars
│   │       └── config.go
│   │
├── migrations/
│   ├── 001_create_addresses.sql
│   └── 002_create_customers.sql
│
├── .gitignore
├── go.mod
└── README.md
```
**Note**: Les autres microservices: product-microservice, order-microservice payment-microservice possèdent le même organigramme.
