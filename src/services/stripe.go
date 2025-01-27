package services

import (
	"fmt"
	"go-fiber-template/src/domain/entities"
	"go-fiber-template/src/domain/repositories"
	"os"
	"time"

	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
)

type StripeService struct {
	UsersRepository repositories.IUsersRepository
}

// NewStripeService creates a new StripeService instance
func NewStripeService(repo0 repositories.IUsersRepository) IStripeService {
	return &StripeService{
		UsersRepository: repo0,
	}
}

type IStripeService interface {
	StripeCreatePrice(userID string, BodyPrice *entities.BodyPrice) (string, error)
	PointIncrease(userId string, point int) error
}

// CreateCheckoutSession creates a Stripe Checkout Session
func (sv StripeService) StripeCreatePrice(userID string, BodyPrice *entities.BodyPrice) (string, error) {
	price := BodyPrice.Price
	currency := BodyPrice.Currency
	method := BodyPrice.Method
	saleCodeName := BodyPrice.Sales
	stripe.Key = os.Getenv("STRIPE_KEY")
	url := os.Getenv("STRIPE_REDIRECT")

	var unitPrice int32
	var err error
	name := fmt.Sprintf("pack %v", price)
	unitPrice = int32(price * 100)
	// packageID := packageData["package_id"].(string)
	var metaData map[string]string
	if BodyPrice.PackageID != "" {
		metaData = map[string]string{"user_id": userID, "action": "point_package", "package_id": BodyPrice.PackageID}
	}

	paymentMethod := method
	if saleCodeName != "" {
		userID = fmt.Sprintf("%v_sid_%v", userID, saleCodeName)
	}
	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice(paymentMethod),

		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String(currency),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String(name),
					},
					UnitAmount: stripe.Int64(int64(unitPrice)),
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode:                stripe.String(string(stripe.CheckoutSessionModePayment)),
		ClientReferenceID:   stripe.String(userID),
		SuccessURL:          stripe.String(url),
		CancelURL:           stripe.String(os.Getenv("FRONT_REDIRECT_URL_STRIPE")),
		AllowPromotionCodes: stripe.Bool(true),
		ExpiresAt:           stripe.Int64(time.Now().Add(60 * time.Minute).Unix()),
		Metadata:            metaData,
	}
	a, err := session.New(params)
	if err != nil {
		return "", err
	}
	return a.URL, nil
}

func (sv StripeService) PointIncrease(userId string, point int) error {
	if err := sv.UsersRepository.UpdatePointStripe(userId,point); err != nil {
		return err
	}
	return nil
}
