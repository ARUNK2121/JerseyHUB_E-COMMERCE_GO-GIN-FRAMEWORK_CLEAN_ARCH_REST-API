package interfaces

import "jerseyhub/pkg/utils/models"

type OfferRepository interface {
	AddNewOffer(model models.OfferMaking) error
	MakeOfferExpire(id int) error
	FindDiscountPercentage(int) (int, error)
}
