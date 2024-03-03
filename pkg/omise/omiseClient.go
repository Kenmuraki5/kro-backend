package omise

import (
	"log"
	"os"

	"github.com/Kenmuraki5/kro-backend.git/domain/restmodel"
	"github.com/joho/godotenv"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Environment variable %s not set", key)
	}
	return value
}

func GetOmiseClient() (*omise.Client, error) {
	publicKey := getEnv("OMISE_PUBLIC_KEY")
	secretKey := getEnv("OMISE_SECRET_KEY")

	client, err := omise.NewClient(publicKey, secretKey)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func CreateSource(client *omise.Client) (*omise.Source, error) {
	result := &omise.Source{}

	err := client.Do(result, &operations.CreateSource{
		Amount:   500000,
		Currency: "thb",
		Type:     "mobile_banking_scb",
	})
	if err != nil {
		return nil, err
	}
	log.Println(result)
	return result, nil
}

func CreateToken(client *omise.Client, payment restmodel.Payment) (string, error) {
	result := &omise.Card{}

	err := client.Do(result, &operations.CreateToken{
		Name:            payment.Name,
		Number:          payment.Number,
		ExpirationMonth: payment.ExpirationMonth,
		ExpirationYear:  payment.ExpirationYear,
		SecurityCode:    payment.Cvc,
	})
	if err != nil {
		return "", err
	}
	log.Println(result)
	return result.ID, nil
}

func CreateChargeBySource(client *omise.Client, sourceID string) error {
	var result omise.Charge

	err := client.Do(&result, &operations.CreateCharge{
		Amount:    500000,
		Currency:  "thb",
		ReturnURI: "http://www.example.com",
		Source:    sourceID,
	})
	if err != nil {
		return err
	}
	log.Println(result.AuthorizeURI)
	return nil
}

func CreateChargeByToken(client *omise.Client, token string, orderId string, total int64) error {
	result := &omise.Charge{}
	err := client.Do(result, &operations.CreateCharge{
		Amount:      total * 100,
		Currency:    "thb",
		Description: "KRO-Gamestore customer charge",
		Card:        token,
		Metadata: map[string]interface{}{
			"OrderID": orderId,
		},
	})
	if err != nil {
		return err
	}
	log.Println(result)
	return nil
}

// func main() {
// 	loadEnv()

// 	client, err := getOmiseClient()
// 	if err != nil {
// 		log.Fatalf("Error creating Omise client: %v", err)
// 	}

// 	// source, err := CreateSource(client)
// 	// if err != nil {
// 	// 	log.Fatalf("Error creating source: %v", err)
// 	// }
// 	// CreateChargeBySource(client, source.ID)

// 	payment := restmodel.Payment{
// 		Name:            "Card Holder",
// 		Number:          "4242424242424242",
// 		ExpirationMonth: 12,
// 		ExpirationYear:  2024,
// 	}
// 	token, err := CreateToken(client, payment)
// 	if err != nil {
// 		log.Fatalf("Error creating token: %v", err)
// 	}
// 	err = CreateChargeByToken(client, token, "123456")
// 	if err != nil {
// 		log.Fatalf("Error creating charge by token: %v", err)
// 	}
// }
