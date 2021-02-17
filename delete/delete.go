package delete

import (
	"fmt"

	"github.com/stripe/stripe-go/v72/product"

	"github.com/javierlopezdeancos/antiqvvs/get"
)

// Start to delete all products
func Start() error {
	fmt.Println("\n🔵 [INFO] Deleting wines...")
	fmt.Println()

	wines, err := get.Products()

	if err != nil {
		fmt.Println(err)
	}

	for _, wine := range wines {
		fmt.Println(wine.ID)
		_, err := product.Del(wine.ID, nil)

		if err != nil {
			return fmt.Errorf("🔴 Error %v deleting wine: %s", err, wine.ID)
		}
	}

	fmt.Println("\n🟢 [SUCCESS] Removed all wines with success, chin! chin! 🍷💢🍷")

	return nil
}
