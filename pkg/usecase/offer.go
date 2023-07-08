package usecase

import (
	interfaces "jerseyhub/pkg/repository/interface"
	"jerseyhub/pkg/utils/models"
)

type offerUseCase struct {
	repository interfaces.OfferRepository
}

func NewOfferUseCase(repo interfaces.OfferRepository) *offerUseCase {
	return &offerUseCase{
		repository: repo,
	}
}

func (off *offerUseCase) AddNewOffer(model models.OfferMaking) error {
	if err := off.repository.AddNewOffer(model); err != nil {
		return err
	}

	return nil
}

func (off *offerUseCase) MakeOfferExpire(id int) error {
	if err := off.repository.MakeOfferExpire(id); err != nil {
		return err
	}

	return nil
}
