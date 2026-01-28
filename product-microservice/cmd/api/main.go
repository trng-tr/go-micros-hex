package main

import (
	"log"

	"github.com/trng-tr/product-microservice/internal/application/ports/in"
	"github.com/trng-tr/product-microservice/internal/application/ports/out"
	"github.com/trng-tr/product-microservice/internal/application/usecase"
	"github.com/trng-tr/product-microservice/internal/infrastructure/config"
	"github.com/trng-tr/product-microservice/internal/infrastructure/in/web/handlers"
	"github.com/trng-tr/product-microservice/internal/infrastructure/in/web/routes"
	"github.com/trng-tr/product-microservice/internal/infrastructure/out/models"
	"github.com/trng-tr/product-microservice/internal/infrastructure/out/repositories"
	"github.com/trng-tr/product-microservice/internal/infrastructure/out/services"
)

func main() {
	var cfg config.AppConfig = config.AppConfig{}
	cfg.LoadConfigurations()
	database, err := cfg.GetDbDNS()
	if err != nil {
		log.Fatal(err)
		return
	}
	var prodRepo services.ProductRepository = repositories.NewProductRepositoryImpl(database)
	var locationRepo services.Repository[models.Location, int64] = repositories.NewLocationRepository(database)
	var stockRepo services.StockRepository = repositories.NewStockRepositoryImpl(database)
	var prodOut out.OutProductService = services.NewOutProductServiceImpl(prodRepo)
	var locationOut out.OutLocationService = services.NewOutLocationServiceImpl(locationRepo)
	var outStk out.OutStockService = services.NewOutOutStockServiceImpl(stockRepo)
	var outUuid out.OutUuidGeneratorService = services.NewOutUuidGeneratorServiceImpl()
	var prodIn in.InProductService = usecase.NewInProductServiceImpl(prodOut, outUuid)
	var locationIn in.InLocationService = usecase.NewInLocationServiceImpl(locationOut)
	var stkIn in.InStockService = usecase.NewInStockServiceImpl(outStk, prodOut)
	var phs routes.ProductHandlerService = handlers.NewProductHandlerServiceImpl(prodIn)
	var lhs routes.LocationHandlerService = handlers.NewLocationHandlerServiceImpl(locationIn)
	var shs routes.StockHanderService = handlers.NewStockHandlerServiceImpl(stkIn, prodIn, locationIn)
	routes := routes.NewRoutes(phs, shs, lhs)
	var engine = routes.RegisterApiRoutes()

	engine.Run(cfg.GetServerAddress())
}
