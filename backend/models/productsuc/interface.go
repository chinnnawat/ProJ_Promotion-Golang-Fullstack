package productsuc

import "github.com/chinnnawat/ProJ_Promotion-Golang/backend/api/domain/entities"

type ProductsDataStorer interface {
	GetAll() []entities.Product
	Create(product *entities.Product) error
}
