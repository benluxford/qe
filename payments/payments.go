package payments

import (
	"encoding/json"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/balance"
)

type Resolver struct {
	credentials credentials
}

type credentials struct {
	stripeKey string
}

func (r *Resolver) stripeBalance() (bal []byte, err error) {
	stripe.Key = r.credentials.stripeKey
	b, err := balance.Get(nil)
	if err != nil {
		return
	}
	bal, err = json.MarshalIndent(b, "", "    ")
	if err != nil {
		panic(err)
	}
	return
}

// NewResolver : return the fresh resolver
func NewResolver(stripeKey string) (r *Resolver) {
	r = new(Resolver)
	// Set the stripe key
	r.credentials.stripeKey = stripeKey
	return
}
