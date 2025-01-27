package entities

// CheckoutSession represents a Stripe Checkout Session
type BodyPrice struct {
	Price    float64  `json:"price" bson:"price,omitempty"`
	Currency string   `json:"currency" bson:"currency,omitempty"`
	Method   []string `json:"method" bson:"method,omitempty"`
	Sales    string   `json:"sales" bson:"sales,omitempty"`

	// optional
	PackageID string `json:"pack_id" bson:"pack_id,omitempty"`
}

