package populate

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

func getWinesFromLocalJSON() ([]wine.Structure, error) {
	var winesJSONFilePath string

	if environment.Config.EnvType == environment.Types["test"] {
		winesJSONFilePath = "./wines/test/wines.json"
	} else if environment.Config.EnvType == environment.Types["production"] {
		winesJSONFilePath = "./wines/prod/wines.json"
	} else {
		panic("the environment TYPE not exist")
	}

	winesData, err := ioutil.ReadFile(winesJSONFilePath)

	if err != nil {
		return nil, err
	}

	var wines []wine.Structure

	err = json.Unmarshal(winesData, &wines)

	if err != nil {
		return nil, err
	}

	return wines, nil
}

func getPricesFromLocalJSON() ([]winePrice.Structure, error) {
	var pricesJSONFilePath string

	if environment.Config.EnvType == environment.Types["test"] {
		pricesJSONFilePath = "./prices/test/prices.json"
	} else if environment.Config.EnvType == environment.Types["production"] {
		pricesJSONFilePath = "./prices/prod/prices.json"
	} else {
		panic("the environment TYPE not exist")
	}

	pricesJSON, err := ioutil.ReadFile(pricesJSONFilePath)

	if err != nil {
		return nil, err
	}

	var prices []winePrice.Structure

	err = json.Unmarshal(pricesJSON, &prices)

	if err != nil {
		return nil, err
	}

	return prices, nil
}

// CreateWines from JSON to stripe
func CreateWines() error {
	fmt.Println("\n🔵 [INFO] Creating wines...")
	fmt.Println()

	wines, err := getWinesFromLocalJSON()

	if err != nil {
		return err
	}

	for _, wine := range wines {
		var images []*string

		for _, i := range wine.Images {
			images = append(images, stripe.String(i))
		}

		paramses := []*stripe.ProductParams{
			{
				ID:     stripe.String(wine.ID),
				Images: images,
				Name:   stripe.String(wine.Name),
				Type:   stripe.String(string(stripe.ProductTypeGood)),
				URL:    stripe.String(wine.URL),
			},
		}

		metadatases := []map[string]string{
			{
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
			},
		}

		for p := 0; p < len(paramses); p++ {
			params := paramses[p]
			metadata := metadatases[p]

			for key, value := range metadata {
				params.AddMetadata(key, value)
			}

			_, err := product.New(params)

			if err != nil {
				stripeErr, ok := err.(*stripe.Error)

				if ok && stripeErr.Code == "resource_already_exists" {
					fmt.Printf("🟡 [WARNING] Wine: %s, already exists", *params.Name)
					fmt.Println()
					fmt.Println()
				} else {
					fmt.Println()
					return fmt.Errorf("🔴 [ERROR] Error getting wine: %s, %v", *params.Name, err)
				}
			}

			fmt.Printf("🟢 [SUCCESS] Created wine: %s\n", *params.Name)
		}
	}

	fmt.Println("\n🟢 [SUCCESS] Your JSON wines are being populated in stripe BBDD")

	return nil
}

// CreatePrices from JSON to stripe
func CreatePrices() error {
	fmt.Println("\n🔵 [INFO] Creating prices...")
	fmt.Println()

	prices, err := getPricesFromLocalJSON()

	if err != nil {
		return err
	}

	for _, prc := range prices {
		paramses := []*stripe.PriceParams{
			{
				Nickname:   stripe.String(prc.Nickname),
				Currency:   stripe.String(string(config.Default().Currency)),
				Product:    stripe.String(prc.Product),
				UnitAmount: stripe.Int64(prc.UnitAmount),
			},
		}

		for p := 0; p < len(paramses); p++ {
			params := paramses[p]

			_, err := price.New(params)

			if err != nil {
				stripeErr, ok := err.(*stripe.Error)

				if ok && stripeErr.Code == "resource_already_exists" {
					fmt.Printf("🟡 [WARNING] Price to wine: %s, already exists", *params.Nickname)
					fmt.Println()
					fmt.Println()
				} else {
					fmt.Println()
					return fmt.Errorf("🔴 [ERROR] Error getting price to wine: %s, %v", *params.Nickname, err)
				}
			}

			fmt.Printf("🟢 [SUCCESS] Created price to wine: %s\n", *params.Nickname)
		}
	}

	fmt.Println("\n🟢 [SUCCESS] Your JSON prices are being populated in stripe BBDD")

	return nil
}

func createProducts() error {
	fmt.Println("\n🔵 [INFO] Creating products...")

	err := CreateWines()

	if err != nil {
		return err
	}

	err = CreatePrices()

	if err != nil {
		return err
	}

	return nil
}

// Start to populate data in stripe
func Start() error {
	err := createProducts()

	if err != nil {
		return err
	}

	fmt.Println("\n🟢 [SUCCESS] Quantvm Stripe eCommerce was populated, chin! chin! 🍷💢🍷")

	return nil
}
