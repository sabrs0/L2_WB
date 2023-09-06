package builder

type Product struct {
	name        string
	description string
	price       float64
}

type ProductBuilder interface {
	SetName(name string)
	SetDescription(description string)
	SetPrice(price float64)
	GetProduct() *Product
}

type LampBuilder struct {
	product Product
}

// дефолтные значения лучше прописывать в конструкторе конкретного билдера
func NewLampBuilder() *LampBuilder {
	lampBuilder := LampBuilder{}
	lampBuilder.SetName("lamp")
	lampBuilder.SetDescription("default 50W")
	lampBuilder.SetPrice(199)

	return &lampBuilder
}

func (lb *LampBuilder) SetName(name string) {
	lb.product.name = name
}
func (lb *LampBuilder) SetDescription(description string) {
	lb.product.description = description
}
func (lb *LampBuilder) SetPrice(price float64) {
	lb.product.price = price
}
func (lb LampBuilder) GetProduct() *Product {
	return &lb.product
}

type TableBuilder struct {
	product Product
}

func NewTableBuilder() *TableBuilder {
	tableBuilder := TableBuilder{}
	tableBuilder.SetName("table")
	tableBuilder.SetDescription("1x1x1")
	tableBuilder.SetPrice(999)
	return &tableBuilder
}

func (tb *TableBuilder) SetName(name string) {
	tb.product.name = name
}
func (tb *TableBuilder) SetDescription(description string) {
	tb.product.description = description
}
func (tb *TableBuilder) SetPrice(price float64) {
	tb.product.price = price
}
func (tb TableBuilder) GetProduct() *Product {
	return &tb.product
}
