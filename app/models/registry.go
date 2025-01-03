package models


type Model struct {
	Model any
}

func RegisterModels() []Model {
	return []Model{
		{Model: User{}},
		{Model: Address{}},
		{Model: Product{}},
		{Model: ProductImage{}},
		{Model: Section{}},
		{Model: Category{}},
		{Model: OrderItem{}},
		{Model: OrderCustomer{}},
		{Model: Payment{}},
		{Model: Shipment{}},
	}
}