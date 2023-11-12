package usecase

import (
	"context"
	"distributed_locking/repository"
)

type PromoUsecase interface {
	GetPromoByCode(ctx context.Context, code string) (resp GetPromoResponse, err error)
	DecrimentPromoByCode(ctx context.Context, code string) (resp string, err error)
}

type usecase struct {
	promoRepository repository.Promo
}

func NewUsecase(promoRepository repository.Promo) *usecase {
	return &usecase{
		promoRepository: promoRepository,
	}
}

func (uc *usecase) GetPromoByCode(ctx context.Context, code string) (resp GetPromoResponse, err error) {
	promo, err := uc.promoRepository.GetPromoByCode(ctx, code)
	if err != nil {
		return
	}

	resp = GetPromoResponse{
		UUID:      promo.UUID,
		Code:      promo.Code,
		Qouta:     promo.Qouta,
		UpdatedAt: promo.UpdatedAt,
	}

	return
}

func (uc *usecase) DecrimentPromoByCode(ctx context.Context, code string) (resp string, err error) {
	err = uc.promoRepository.DecrimentPromoByCode2(ctx, code)
	if err != nil {
		return
	}

	resp = "success claim promo code " + code

	return
}
