package services

/*en Go, les interfaces sont du côte de celui qui les utilise(DI)
et non pas du coté de cui qui les implémente
*/
import "github.com/trng-tr/customer-microservice/internal/infrastructure/out/models"

/* CustomerRepository interface extends  interfaces:
XXXSaveAndReadOneORepository,XXXReadAllORepository,
XXXUpdateORepository, XXXDeleteOneORepository, methods:
SaveO, FindOByID,FindAllO,UpdateO,DeleteO
*/
type CustomerRepository interface {
	XXXSaveAndReadOneORepository[models.CustomerModel, int64]
	XXXReadAllORepository[models.CustomerModel]
	XXXUpdateORepository[models.CustomerModel, int64]
	XXXDeleteOneORepository[models.CustomerModel, int64]
}

/* AddressRepository interface extends  interfaces:
XXXSaveAndReadOneORepository,XXXReadAllORepository,
XXXDeleteOneORepository: methods SaveO, FindOByID,
FindAllO& DeleteO
*/
type AddressRepository interface {
	XXXSaveAndReadOneORepository[models.AddressModel, int64]
	XXXReadAllORepository[models.AddressModel]
	XXXDeleteOneORepository[models.CustomerModel, int64]
}
