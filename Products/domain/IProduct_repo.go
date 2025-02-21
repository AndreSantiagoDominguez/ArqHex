package domain

type Iproduct interface {
	Save(product *Product) (uint, error);
	GetAllProduct() []Product;
	EditProduct(id int ,product Product) (uint,error)
	DeleteProduct(id int) (uint, error)
}