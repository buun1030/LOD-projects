package workshop_test

import (
	"testing"

	"github.com/NuttapolCha/lod-workshop-go/workshop"
	"github.com/stretchr/testify/assert"
)

func NewTestingPromotion() workshop.Promotion {
	return workshop.Promotion{
		Discount: 10,
	}
}

func TestHandlePersonPurchaseProductsWithPromotion(t *testing.T) {
	prmotion := NewTestingPromotion()

	t.Run("person with enough money buying 1 quantity of product", func(t *testing.T) {
		initialPerson := NewTestingPerson()
		initialShop := NewTestingShop()

		buyingProduct := "bread"
		expectedPersonAfterBuying := &workshop.Person{
			Name: "Youth",
			Wallet: workshop.Wallet{
				Cash: workshop.Money{
					Amount: 60,
				},
			},
		}
		expectedShopAfterBuying := &workshop.Shop{
			Name: "7-11",
			Shelves: []workshop.Shelf{
				{
					Name: "food",
					Products: []workshop.Product{
						{
							Name: "bread",
							Price: workshop.Money{
								Amount: 50,
							},
							Quantity: 1,
						},
						{
							Name: "egg",
							Price: workshop.Money{
								Amount: 20,
							},
							Quantity: 2,
						},
					},
				},
				{
					Name: "drink",
					Products: []workshop.Product{
						{
							Name: "water",
							Price: workshop.Money{
								Amount: 10,
							},
							Quantity: 2,
						},
						{
							Name: "milk",
							Price: workshop.Money{
								Amount: 20,
							},
							Quantity: 2,
						},
						{
							Name: "whey protein",
							Price: workshop.Money{
								Amount: 200,
							},
							Quantity: 2,
						},
					},
				},
			},
		}
		expectedCanPurchase := true

		actualCanPurchase := workshop.HandlePersonPurchaseProductsWithPromotion(initialPerson, initialShop, prmotion, buyingProduct, 1)
		assert.Equal(t, expectedPersonAfterBuying, initialPerson)
		assert.Equal(t, expectedShopAfterBuying, initialShop)
		assert.Equal(t, expectedCanPurchase, actualCanPurchase)
	})

	t.Run("person with enough money buying 2 quantity of product", func(t *testing.T) {
		initialPerson := NewTestingPerson()
		initialShop := NewTestingShop()

		buyingProduct := "egg"
		expectedPersonAfterBuying := &workshop.Person{
			Name: "Youth",
			Wallet: workshop.Wallet{
				Cash: workshop.Money{
					Amount: 70,
				},
			},
		}
		expectedShopAfterBuying := &workshop.Shop{
			Name: "7-11",
			Shelves: []workshop.Shelf{
				{
					Name: "food",
					Products: []workshop.Product{
						{
							Name: "bread",
							Price: workshop.Money{
								Amount: 50,
							},
							Quantity: 2,
						},
						{
							Name: "egg",
							Price: workshop.Money{
								Amount: 20,
							},
							Quantity: 0,
						},
					},
				},
				{
					Name: "drink",
					Products: []workshop.Product{
						{
							Name: "water",
							Price: workshop.Money{
								Amount: 10,
							},
							Quantity: 2,
						},
						{
							Name: "milk",
							Price: workshop.Money{
								Amount: 20,
							},
							Quantity: 2,
						},
						{
							Name: "whey protein",
							Price: workshop.Money{
								Amount: 200,
							},
							Quantity: 2,
						},
					},
				},
			},
		}
		expectedCanPurchase := true

		actualCanPurchase := workshop.HandlePersonPurchaseProductsWithPromotion(initialPerson, initialShop, prmotion, buyingProduct, 2)
		assert.Equal(t, expectedPersonAfterBuying, initialPerson)
		assert.Equal(t, expectedShopAfterBuying, initialShop)
		assert.Equal(t, expectedCanPurchase, actualCanPurchase)
	})

	t.Run("person with not enough money", func(t *testing.T) {
		initialPerson := NewTestingPerson()
		initialPersonValue := *initialPerson
		initialShop := NewTestingShop()
		initialShopValue := *initialShop

		buyingProduct := "whey protein"
		expectedPersonAfterBuying := &initialPersonValue
		expectedShopAfterBuying := &initialShopValue
		expectedCanPurchase := false

		actualCanPurchase := workshop.HandlePersonPurchaseProductsWithPromotion(initialPerson, initialShop, prmotion, buyingProduct, 1)
		assert.Equal(t, expectedPersonAfterBuying, initialPerson)
		assert.Equal(t, expectedShopAfterBuying, initialShop)
		assert.Equal(t, expectedCanPurchase, actualCanPurchase)
	})

	t.Run("person buying product that is not in the shop", func(t *testing.T) {
		initialPerson := NewTestingPerson()
		initialPersonValue := *initialPerson
		initialShop := NewTestingShop()
		initialShopValue := *initialShop

		buyingProduct := "bazooka"
		expectedPersonAfterBuying := &initialPersonValue
		expectedShopAfterBuying := &initialShopValue
		expectedCanPurchase := false

		actualCanPurchase := workshop.HandlePersonPurchaseProductsWithPromotion(initialPerson, initialShop, prmotion, buyingProduct, 1)
		assert.Equal(t, expectedPersonAfterBuying, initialPerson)
		assert.Equal(t, expectedShopAfterBuying, initialShop)
		assert.Equal(t, expectedCanPurchase, actualCanPurchase)
	})

	t.Run("not enough products in stock", func(t *testing.T) {
		initialPerson := NewTestingPerson()
		initialPersonValue := *initialPerson
		initialShop := NewTestingShop()
		initialShopValue := *initialShop

		buyingProduct := "egg"
		expectedPersonAfterBuying := &initialPersonValue
		expectedShopAfterBuying := &initialShopValue
		expectedCanPurchase := false

		actualCanPurchase := workshop.HandlePersonPurchaseProductsWithPromotion(initialPerson, initialShop, prmotion, buyingProduct, 3)
		assert.Equal(t, expectedPersonAfterBuying, initialPerson)
		assert.Equal(t, expectedShopAfterBuying, initialShop)
		assert.Equal(t, expectedCanPurchase, actualCanPurchase)
	})
}
