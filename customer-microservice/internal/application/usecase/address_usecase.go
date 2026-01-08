package usecase

import "github.com/trng-tr/customer-microservice/internal/application/ports/out"

/*AddressServiceImpl implement port d'entrée exposé à l'extreieur
il utilise pour cela le port de sortie OutAddressService*/
type AddressServiceImpl struct {
	OutService out.OutAddressService
}
