package domain

type Product struct {
	ID    int64
	Name  string
	Price int64
}

func (p Product) GetAllProduct() any {
	panic("unimplemented")
}

func NewProduct(name string, price int64) *Product {
	return &Product{
		ID:    1,
		Name:  name,
		Price: int64(price),
	}
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) SetName(name string) {
	p.Name = name
}
