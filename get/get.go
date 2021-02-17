package get

import (
	"fmt"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/price"
	"github.com/stripe/stripe-go/v72/product"
)

// Products get all products
func Products() ([]*stripe.Product, error) {
	fmt.Println("\n🔵 [INFO] Get all wines to remove them...")
	fmt.Println()

	products := []*stripe.Product{}

	plp := &stripe.ProductListParams{}
	plp.Filters.AddFilter("limit", "", "3")

	l := product.List(plp)

	fmt.Println()

	for l.Next() {
		fmt.Printf("🟢 [SUCCESS] Getting wine:  %s\n", l.Product().ID)
		products = append(products, l.Product())
	}

	err := l.Err()

	if err != nil {
		return nil, fmt.Errorf("🔴 [ERROR] Error getting products: %v", err)
	}

	fmt.Println("\n🟢 [SUCCESS] You got now all wines, chin! chin!, 🍷💢🍷")

	return products, nil
}

// Prices get all products
func Prices() ([]*stripe.Price, error) {
	fmt.Println("\n🔵 [INFO] Get all prices to remove them...")

	prices := []*stripe.Price{}

	plp := &stripe.PriceListParams{}
	plp.Filters.AddFilter("limit", "", "3")

	l := price.List(plp)

	for l.Next() {
		fmt.Printf("🟢 [SUCCESS] Getting price:  %s\n", l.Price().ID)
		prices = append(prices, l.Price())
	}

	err := l.Err()

	if err != nil {
		return nil, fmt.Errorf("🔴 [ERROR] Error getting prices: %v", err)
	}

	fmt.Println("\n🟢 [SUCCESS] You got now all prices, chin! chin! 🍷💢🍷")

	return prices, nil
}
