# HexaShop : Microservices + Archi Hexagonale

# 1. Présentation générale

Le projet **go-micros-hex** est un écosystème de microservices écrit en **Go**, structuré selon les principes de l’**architecture hexagonale (Ports & Adapters)**.  
Il illustre une plateforme de type e-commerce modulaire, avec séparation claire des responsabilités métier, techniques et d’orchestration.

L’architecture repose sur :
- des microservices indépendants
- une base de données par microservice
- une **API Gateway Kong** comme point d’entrée unique
- des échanges synchrones REST entre services

---

# 2. Fonction principale (vision système)

La fonction principale du système est de **gérer le cycle de vie d’une commande client**, depuis la validation des produits et du stock jusqu’à la persistance de la commande, dans un environnement distribué.

Le système permet :
- la gestion des clients
- la gestion des produits, stocks et localisations
- la création de commandes avec vérification du stock
- la mise à jour des quantités disponibles

---

# 3. Architecture globale

## 3.1 Microservices

| Microservice | Port | Responsabilité principale 
|--------------|------|---------------------------
| customer-microservice | 8081 | Gestion des clients 
| product-microservice  | 8082 | Produits, stocks, locations
| order-microservice    | 8083 | Commandes & orchestration
| Kong API Gateway      | 8080 | Point d’entrée unique

Chaque microservice :
- est autonome
- possède sa propre base de données
- expose une API REST
- implémente une architecture hexagonale

---

## 3.2 Architecture hexagonale (Ports & Adapters)

A l'instar de customer-microservice, chaque microservice est structuré de la manière suivante :
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
│   │   │       ├── fieds_checker.go                        # validation des champs du métier
│   │   |       └── errors.go                               # erreurs métier
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
│   │   ├── out/                                            # ✅ save dans la db
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
│   |   ├── config/                                         # 4️⃣ la config des env vars
│   │       └── config.go
│   │
├── migrations/                                             
│   └── 001_create_tables.sql                               
│
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```
Les autres microservices: product-microservice, order-microservice possèdent le même organigramme.

## Architecture Microservices

![Architecture HexaShop](microservices.png)

LÉGENDE
-------
- Chaque microservice est autonome et possède sa propre base de données
- Les communications inter-services se font via HTTP REST
- OrderSvc orchestre les appels vers les microservices CustomerSvc et ProductSvc
- Kong Api Gateway agit comme point d’entrée unique (API Gateway)
- Architecture Hexagonale basée sur les principes Ports & Adapters


Le domaine ne dépend jamais :
- ni du HTTP
- ni de la base de données
- ni des autres microservices

---

# 4. API Gateway – Kong

## 4.1 Rôle de Kong

**Kong** agit comme **API Gateway** et constitue le **point d’entrée unique** du système.

Il est responsable de :
- router les requêtes HTTP vers les microservices
- centraliser l’accès aux APIs
- préparer l’ajout de préoccupations transverses :
    - caching,
    - rate limiting,
    - logs,
    - CORS

La configuration est déclarative via le fichier `./kong.yml`

---

## 4.2 Flux avec Kong

Client
    ↓
Kong API Gateway
    ↓
Microservice cible (Customer / Product / Order)

Exemple :  

POST /api/v1/orders
→ Kong
→ order-microservice
→ product-microservice (stock, location, produit)
→ base de donnée

---

# 5. Documentation par microservice

---

## 5.1 Customer Microservice (8081)

### Fonction principale
Gérer les **clients** et leurs informations associées.

### Périmètre fonctionnel
- Création d’un client
- Consultation d’un client
- Mise à jour des informations
- Suppression ou désactivation logique

### Responsabilités
- Validation des données client
- Persistance en base
- Exposition d’une API REST

### Exemples d’API
- GET /api/v1/customers
- GET /api/v1/customers/{id}
- POST /api/v1/customers
- PUT /api/v1/customers/{id}
- DELETE /api/v1/customers/{id}

---

## 5.2 Product Microservice (8082)

### Fonction principale
Gérer le **catalogue produit**, les **stocks** et les **localisations**.

### Périmètre fonctionnel

#### Produits
- Création / mise à jour
- Activation / désactivation
- Consultation

#### Locations
- Gestion des lieux (villes / entrepôts)

#### Stocks
- Stock par couple `(product, location)`
- Quantité toujours ≥ 0
- Consultation et mise à jour

### APIs principales
- POST /api/v1/locations
- POST /api/v1/products
- POST /api/v1/stocks
- GET /api/v1/products/{id}
- GET /api/v1/locations/{id}
- GET /api/v1/stocks
- GET /api/v1/stocks/locations/{locationId}/products/{productId}
- PUT /api/v1/stocks/locations/{locationId}/products/{productId}/set-qte  
Body json attendu pour la mise à jour (PUT) :
```
{
  "quantity": 150
}
```

## 5.3 Order Microservice (8082)

### Fonction principale
Gérer les **commandes** et orchestrer les appels vers les autres microservices.

### Périmètre fonctionnel

- Création d’une commande
- Gestion des lignes de commande
- Vérification produit & stock
- Mise à jour du stock distant
- Persistance de la commande

### Flux de création d’une commande
1. Réception de la requête client
2. Vérification du client et sont state (remote Customer microservice)
2. Vérification des produits et leur state (remote Product microservice)
3. Vérification du stock par ligne (remote Product microservice)
4. Refus si stock insuffisant (remote Product microservice)
5. Mise à jour du stock (remote Product microservice)
6. Enregistrement de la commande

#### Exemple de requête
```
{
  "customer_id": 1,
  "order_lines": [
    { "product_id": 2, "location_id": 1, "quantity": 5 },
    { "product_id": 1, "location_id": 1, "quantity": 3 }
  ]
}
```

# 6. Documentation développeur

## Pré-requis
- Go (modules)
- Docker & Docker Compose
- PostgreSQL

## Démarrage
- Récupérer le repo git `git clone -b main https://github.com/trng-tr/go-micros-hex.git`
- Rentrer à la racine et builder pour construire `docker compose -f stack-docker.yaml build`
- Lancer la stack via stack-docker.yml `docker compose -f stack-docker.yaml up -d`
- Vérifier la disponibilité de Kong et des microservices
- Utiliser Kong comme point d’entrée principal

## Variables d’environnement (exemple)
- APP_HOSTNAME 
- PORT
- DB_HOST
- DB_PORT
- DB_NAME
- DB_USER
- DB_PASSWORD
- PRODUCT_SERVICE_URL (Order Microservice)
- CUSTOMER_SERVICE_URL (Order Microservice)

# 7. Gestion des erreurs

Les erreurs sont renvoyées sous forme :
```
{
  "status": "FAIL",
  "message": "description",
  "created_at": "ISO_DATE"
}
```