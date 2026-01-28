package services

/*en Go, les interfaces sont du côte de celui qui les utilise(DI)
et non pas du coté de cui qui les implémente, ici c'est OutProductServiceImpl
 et OutStockServiceImpl qui les injecte par DI
*/
import "context"

// Repository common methods for Product and stock repos
type Repository[O any, ID comparable] interface {
	Save(ctx context.Context, o O) (O, error)
	FindAll(ctx context.Context) ([]O, error)
	FindByID(ctx context.Context, id ID) (O, error)
}
