package main

import (
	"log"

	"github.com/trng-tr/order-microservice/internal/application/in"
	"github.com/trng-tr/order-microservice/internal/application/out"
	"github.com/trng-tr/order-microservice/internal/application/usecase"
	"github.com/trng-tr/order-microservice/internal/infrastructure/config"
	"github.com/trng-tr/order-microservice/internal/infrastructure/in/http/handlers"
	"github.com/trng-tr/order-microservice/internal/infrastructure/in/http/routes"
	"github.com/trng-tr/order-microservice/internal/infrastructure/out/repositories"
	"github.com/trng-tr/order-microservice/internal/infrastructure/out/services"
)

func main() {
	var appConf config.AppConfig = config.AppConfig{}
	appConf.LoadConfig()
	postgresDb, err := appConf.GetDbServer()
	if err != nil {
		log.Fatal(err)
		return
	}
	var orderRepo services.OrderRepo = repositories.NewOrderRepoImpl(postgresDb)
	var orderLineRepo services.OrderLineRepo = repositories.NewOrderLineRepoImpl(postgresDb)
	var outOrderSvc out.OutOrderService = services.NewOutOrderServiceImpl(orderRepo)
	var outOrderLineSvc out.OutOrderLineService = services.NewOutOrderLineServiceImpl(orderLineRepo)
	var remoteProductSvc out.RemoteProductService = services.NewRemoteProductServiceImpl(appConf.ProductBaseUrl)
	var outRemoteCustomerSvc out.RemoteCustomerService = services.NewRemoteCustomerServiceImpl(appConf.CustomerBaseUrl)
	var inOrderSvc in.InOrderService = usecase.NewOrderUseCase(outOrderSvc, outRemoteCustomerSvc, remoteProductSvc)
	var inOrderLineSvc in.InOrderLineService = usecase.NewOrderLineUseCase(outOrderLineSvc, outOrderSvc, remoteProductSvc)
	var inRemoteCustomer in.RemoteCustomerService = usecase.NewRemoteCustomerServiceImpl(outRemoteCustomerSvc)
	var inRemoteProduct in.RemoteProductService = usecase.NewRemoteProductServiceImpl(remoteProductSvc)
	var handler routes.OrderHandler = handlers.NewOrderHandlerImpl(inOrderSvc, inOrderLineSvc, inRemoteCustomer, inRemoteProduct)
	var routeRegistration = routes.NewRouteRegistration(handler)
	var engine = routeRegistration.RegisterRoutes()
	engine.Run(appConf.GetAppServer())
}
