package interfaces

import (
	"jerseyhub/pkg/domain"
	"jerseyhub/pkg/utils/models"
)

type OfferUseCase interface {
	AddNewOffer(model models.OfferMaking) error
	MakeOfferExpire(id int) error
	GetOffers() ([]domain.Offer, error)
}
