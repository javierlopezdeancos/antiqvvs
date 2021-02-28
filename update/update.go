package update

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/price"
	"github.com/stripe/stripe-go/v72/product"

	"github.com/javierlopezdeancos/antiqvvs/config"
	"github.com/javierlopezdeancos/antiqvvs/environment"
	winePrice "github.com/javierlopezdeancos/antiqvvs/price"
	"github.com/javierlopezdeancos/antiqvvs/wine"
)

func getWineFromLocalJSON() (wine.Structure, error) {
	var wineJSONFilePath string
	var wine wine.Structure

	if environment.Config.EnvType == environment.Types["test"] {
		wineJSONFilePath = "./wine/test/wine.json"
	} else if environment.Config.EnvType == environment.Types["production"] {
		wineJSONFilePath = "./wine/prod/wine.json"
	} else {
		panic("the environment TYPE not exist")
	}

	wineJSON, err := ioutil.ReadFile(wineJSONFilePath)

	if err != nil {
		return wine, err
	}

	err = json.Unmarshal(wineJSON, &wine)

	if err != nil {
		return wine, err
	}

	return wine, nil
}

func getPriceFromLocalJSON() (winePrice.Structure, error) {
	var priceJSONFilePath string
	var price winePrice.Structure

	if environment.Config.EnvType == environment.Types["test"] {
		priceJSONFilePath = "./price/test/price.json"
	} else if environment.Config.EnvType == environment.Types["production"] {
		priceJSONFilePath = "./price/prod/price.json"
	} else {
		panic("the environment TYPE not exist")
	}

	priceJSON, err := ioutil.ReadFile(priceJSONFilePath)

	if err != nil {
		return price, err
	}

	err = json.Unmarshal(priceJSON, &price)

	if err != nil {
		return price, err
	}

	return price, nil
}

// Wine update wine from JSON to stripe
func Wine() error {
	fmt.Println("\n🔵 [INFO] Updating wine...")
	fmt.Println()

	wine, err := getWineFromLocalJSON()

	if err != nil {
		return err
	}

	var images []*string

	for _, i := range wine.Images {
		images = append(images, stripe.String(i))
	}

	params := &stripe.ProductParams{
		Images: images,
		Name:   stripe.String(wine.Name),
		URL:    stripe.String(wine.URL),
	}

	metadata := map[string]string{
		"barrel":           wine.Metadata.Barrel,
		"brandImage":       wine.Metadata.BrandImage,
		"capacity":         wine.Metadata.Capacity,
		"cellar":           wine.Metadata.Cellar,
		"cellarUrl":        wine.Metadata.CellarURL,
		"color":            wine.Metadata.Color,
		"cork":             wine.Metadata.Cork,
		"do":               wine.Metadata.Do,
		"doImage":          wine.Metadata.DoImage,
		"graduation":       wine.Metadata.Graduation,
		"grape":            wine.Metadata.Grape,
		"placeholderImage": wine.Metadata.PlaceholderImage,
		"path":             wine.Metadata.Path,
		"where":            wine.Metadata.Where,
	}

	for key, value := range metadata {
		params.AddMetadata(key, value)
	}

	_, err = product.Update(
		wine.ID,
		params,
	)

	if err != nil {
		stripeErr, _ := err.(*stripe.Error)
		return fmt.Errorf("🔴 [ERROR] Error updating wine: %s, %v", *params.Name, stripeErr)
	}

	fmt.Printf("🟢 [SUCCESS] Update wine: %s\n", *params.Name)

	return nil
}

// Price from JSON to stripe
func Price() error {
	fmt.Println("\n🔵 [INFO] Updating price...")
	fmt.Println()

	prc, err := getPriceFromLocalJSON()

	if err != nil {
		return err
	}

	params := &stripe.PriceParams{
		Nickname:   stripe.String(prc.Nickname),
		Currency:   stripe.String(string(config.Default().Currency)),
		Product:    stripe.String(prc.Product),
		UnitAmount: stripe.Int64(prc.UnitAmount),
	}

	_, err = price.Update(
		prc.ID,
		params,
	)

	if err != nil {
		stripeErr, _ := err.(*stripe.Error)
		return fmt.Errorf("🔴 [ERROR] Error updating price to wine: %s, %v", *params.Nickname, stripeErr)
	}

	fmt.Printf("🟢 [SUCCESS] Update price to wine: %s\n", *params.Nickname)

	return nil
}

func updateProducts() error {
	fmt.Println("\n🔵 [INFO] Updating products...")

	err := Wine()

	if err != nil {
		return err
	}

	err = Price()

	if err != nil {
		return err
	}

	return nil
}

// Start to populate data in stripe
func Start() error {
	err := updateProducts()

	if err != nil {
		return err
	}

	fmt.Println("\n🟢 [SUCCESS] Quantvm Stripe eCommerce was update, chin! chin! 🍷💢🍷")

	return nil
}
