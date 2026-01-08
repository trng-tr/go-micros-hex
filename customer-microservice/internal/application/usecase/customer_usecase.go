package usecase

import "github.com/trng-tr/customer-microservice/internal/application/ports/out"

/*CustomerServiceImpl implement port d'entrée exposé à l'extreieur
il utilise pour cela le port de sortie OutCustomerService*/
type CustomerServiceImpl struct {
	OutService out.OutCustomerService
}
