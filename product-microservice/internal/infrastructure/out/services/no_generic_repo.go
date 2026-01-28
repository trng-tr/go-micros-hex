package services

/*en Go, les interfaces sont du côte de celui qui les utilise(DI)
et non pas du coté de cui qui les implémente, ici c'est OutProductServiceImpl
 et OutStockServiceImpl qui les injecte par DI
*/
import (
	"context"

	"github.com/trng-tr/product-microservice/internal/infrastructure/out/models"
)

// ProductRepository interface extends Repository interface
type ProductRepository interface {
	Repository[models.ProductModel, int64] //extend
	//other methods for contract 👇
	PatchProduct(ctx context.Context, id int64, o models.ProductModel) (models.ProductModel, error)
	DeleteProduct(ctx context.Context, id int64) error
	FindProductBySku(ctx context.Context, sku string) (models.ProductModel, error)
}

// StockRepository interface extends Repository interface
type StockRepository interface {
	Repository[models.StockModel, int64] //extend
	//other method for contract 👇
	UpdateStockQuantity(ctx context.Context, productID int64, quantity int64) (models.StockModel, error)
	FindStockByProductID(ctx context.Context, productID int64) (models.StockModel, error)
}
