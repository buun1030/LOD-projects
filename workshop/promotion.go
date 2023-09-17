package workshop

type Promotion struct {
	Discount int
}

func HandlePersonPurchaseProductsWithPromotion(person *Person, shop *Shop, promotion Promotion, productName string, quantity int) bool {
	if !shop.HasEnoughQuantity(productName, quantity) {
		return false
	}
	totalPrice := shop.CalculateTotalPrice(productName, quantity) - promotion.Discount
	if !person.CanAfford(totalPrice) {
		return false
	}

	person.Pay(totalPrice)           // update person's wallet
	shop.Sell(productName, quantity) // update shop's shelf

	return true // return true if person can afford the product
}
