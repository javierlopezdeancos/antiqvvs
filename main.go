package main

import (
	"fmt"
	"os"
	"path"

	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/price"
	"github.com/stripe/stripe-go/v72/product"

	"github.com/javierlopezdeancos/antiqvvs/config"
)

func loadEnv() error {
	err := godotenv.Load(path.Join("./", ".env"))

	if err != nil {
		return err
	}

	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	if stripe.Key == "" {
		panic("STRIPE_SECRET_KEY must be in environment")
	}

	return nil
}

func createData() error {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	if stripe.Key == "" {
		panic("STRIPE_SECRET_KEY must be in environment")
	}

	err := createProducts()

	if err != nil {
		return err
	}

	err = createPrices()

	if err != nil {
		return err
	}

	return nil
}

func createProducts() error {
	paramses := []*stripe.ProductParams{
		{
			ID: stripe.String("product-wine-bottle-75cl-cristal-3-ases-hocicon"),
			Images: []*string{
				stripe.String("hocicon-big.png"),
				stripe.String("hocicon-medium.png"),
				stripe.String("hocicon-small.png"),
			},
			Name: stripe.String("Hocicón"),
			Type: stripe.String(string(stripe.ProductTypeGood)),
			URL:  stripe.String("https://www.quantvm.es/vinos/3ases/hocicon"),
		},
	}

	metadatases := []map[string]string{
		{
			"barrel":      "8 meses de crianza sobre lías",
			"brandImage":  "brand.jpg",
			"capacity":    "75cl",
			"cellar":      "3 Ases",
			"cellarUrl":   "https://www.3asesvino.com",
			"color":       "Rosado",
			"cork":        "Cristal",
			"do":          "Ribera de Duero",
			"doImage":     "ribera-duero.jpg",
			"graduation":  "13,5º",
			"grape":       "100% Tempranillo",
			"placeholder": "placeholder.png",
			"name":        "Hocicón",
			"path":        "/vinos/3ases/hocicon",
			"where":       "Ribera de Duero",
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
				// This is fine — we expect this to be idempotent.
			} else {
				return err
			}
		}
	}

	return nil
}

func createPrices() error {
	params := &stripe.PriceParams{
		Nickname:   stripe.String("price-wine-bottle-75cl-cristal-3-ases-hocicon"),
		Currency:   stripe.String(string(config.Default().Currency)),
		Product:    stripe.String("product-wine-bottle-75cl-cristal-3-ases-hocicon"),
		UnitAmount: stripe.Int64(1074),
	}

	_, err := price.New(params)

	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := loadEnv()

	if err != nil {
		panic(fmt.Sprintf("error loading .env: %v", err))
	}

	err = createData()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Quantvm Stripe eCommerce was configure")
}
