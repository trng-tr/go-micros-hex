package services

/*en Go, les interfaces sont du côte de celui qui les utilise(DI)
et non pas du coté de cui qui les implémente
*/
import "context"

//XXXSaveAndReadOneORepository generic interface
type XXXSaveAndReadOneORepository[O any, ID comparable] interface {
	SaveO(ctx context.Context, o O) (O, error)
	FindOByID(ctx context.Context, id ID) (O, error)
}

//XXXReaderAllORepository generic interface
type XXXReadAllORepository[O any] interface {
	FindAllO(ctx context.Context) ([]O, error)
}

//XXXUpdateORepository generic interface
type XXXUpdateORepository[O any, ID comparable] interface {
	UpdateO(ctx context.Context, id ID, o O) (O, error)
}

//XXXDeleterOneORepository generic interface
type XXXDeleteOneORepository[O any, ID comparable] interface {
	DeleteO(ctx context.Context, id ID) error
}
