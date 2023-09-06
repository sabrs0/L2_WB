package builder

type ProductDirector struct {
	builder ProductBuilder
}

func NewProductDirector() *ProductDirector {
	return &ProductDirector{}
}
func (d *ProductDirector) SetBuilder(builder ProductBuilder) {
	d.builder = builder
}

func (d ProductDirector) Construct() Product {
	return *d.builder.GetProduct()
}
