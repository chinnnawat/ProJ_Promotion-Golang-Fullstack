package productsuc

import (
	"context"

	"github.com/chinnnawat/ProJ_Promotion-Golang/backend/api/domain/entities"
)

type getProductsUseCase struct {
	dataStore ProductsDataStorer
}

func NewGetProductsUseCase(ds ProductsDataStorer) *getProductsUseCase {
	return &getProductsUseCase{
		dataStore: ds,
	}
}

func (uc *getProductsUseCase) GetProducts(ctx context.Context) []entities.Product {
	all := uc.dataStore.GetAll()
	return all
}
