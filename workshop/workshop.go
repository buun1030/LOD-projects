package workshop

type Person struct {
	Name   string
	Wallet Wallet
}

func (person Person) CanAfford(amount int) bool { //
	return person.Wallet.Cash.Amount >= amount
}

func (person *Person) Pay(amount int) { //
	person.Wallet.Cash.Amount -= amount
}

type Wallet struct {
	Cash Money
}

type Money struct {
	Amount int
}

type Shop struct {
	Name    string
	Shelves []Shelf
}

func (s Shop) HasEnoughQuantity(productName string, quantity int) bool { //
	for _, shelf := range s.Shelves {
		product := shelf.GetProduct(productName)
		if product != nil && product.Quantity >= quantity {
			return true
		}
	}
	return false
}

func (s Shop) CalculateTotalPrice(productName string, quantity int) int { //
	for _, shelf := range s.Shelves {
		totalPrice := shelf.CalculateTotalPrice(productName, quantity)
		if totalPrice > 0 {
			return totalPrice
		}
	}
	return 0
}

func (s *Shop) Sell(productName string, quantity int) { //
	for _, shelf := range s.Shelves {
		if shelf.Sell(productName, quantity) {
			return
		}
	}
	return
}

func (s *Shelf) Sell(productName string, quantity int) bool { //
	for i, product := range s.Products {
		if product.Name == productName && product.Quantity >= quantity {
			s.Products[i].ReduceQuantity(quantity)
			return true
		}
	}
	return false
}

type Shelf struct {
	Name     string
	Products []Product
}

func (s Shelf) GetProduct(productName string) *Product { //
	for _, product := range s.Products {
		if product.Name == productName {
			return &product
		}
	}
	return nil
}

func (s Shelf) CalculateTotalPrice(productName string, quantity int) int { //
	product := s.GetProduct(productName)
	if product != nil {
		return product.CalculateTotalPrice(quantity)
	}
	return 0
}

type Product struct {
	Name     string
	Price    Money
	Quantity int
}

func (p Product) CalculateTotalPrice(quantity int) int { //
	return p.Price.Amount * quantity
}

func (p *Product) ReduceQuantity(quantity int) {
	p.Quantity -= quantity
}

func HandlePersonPurchaseProducts(person *Person, shop *Shop, productName string, quantity int) bool {
	if !shop.HasEnoughQuantity(productName, quantity) {
		return false
	}
	totalPrice := shop.CalculateTotalPrice(productName, quantity)
	if !person.CanAfford(totalPrice) {
		return false
	}

	person.Pay(totalPrice)           // update person's wallet
	shop.Sell(productName, quantity) // update shop's shelf

	return true // return true if person can afford the product
}
