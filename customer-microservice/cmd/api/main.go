package main

import (
	"log"

	_ "github.com/lib/pq" //db connector
	"github.com/trng-tr/customer-microservice/internal/application/usecase"
	"github.com/trng-tr/customer-microservice/internal/infrastructure/config"
	"github.com/trng-tr/customer-microservice/internal/infrastructure/in/http/handlers"
	"github.com/trng-tr/customer-microservice/internal/infrastructure/in/http/routes"
	"github.com/trng-tr/customer-microservice/internal/infrastructure/out/repositories/impl"
	"github.com/trng-tr/customer-microservice/internal/infrastructure/out/services"
)

func main() {
	var cfg config.AppConfig = config.AppConfig{}
	cfg.LoadConfig()
	db, err := cfg.GetDBDns()
	if err != nil {
		log.Fatal(err)
		return
	}
	var addressRepo = impl.NewAddressRepositoryImpl(db)
	var customerRepo = impl.NewCustomerRepositoryImpl(db)
	var addressOutPort = services.NewOutAddressServiceImpl(addressRepo)
	var customerOutPort = services.NewOutCustomerServiceImpl(customerRepo)
	var addressInPort = usecase.NewInAddressServiceImpl(addressOutPort)
	var customerInPort = usecase.NewInCustomerServiceImpl(customerOutPort, addressInPort)
	var customerHandler = handlers.NewCustomerHandlerServiceImpl(customerInPort, addressInPort)
	var addressHandler = handlers.NewAddressHandlerServiceImpl(addressInPort)
	var rr = routes.NewRoutesRegistration(addressHandler, customerHandler)
	var engine = rr.RegisterRoutes()
	engine.Run(cfg.GetAppServer())
}
