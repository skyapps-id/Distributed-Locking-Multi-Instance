package handler

import (
	"context"
	"distributed_locking/usecase"
	"distributed_locking/util"
	"errors"
	"net/http"
)

type PromoHandler interface {
	GetPromoByCode(w http.ResponseWriter, req *http.Request)
}

type handler struct {
	promoHandler usecase.PromoUsecase
}

func NewHandler(promoHandler usecase.PromoUsecase) *handler {
	return &handler{
		promoHandler: promoHandler,
	}
}

func (h *handler) GetPromoByCode(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		var ctx = context.Background()
		code := req.URL.Query().Get("code")
		result, err := h.promoHandler.GetPromoByCode(ctx, code)

		util.Response(w, result, err)
	} else {
		util.Response(w, nil, errors.New("error method"))
	}
}

func (h *handler) DecrimentPromoByCode(w http.ResponseWriter, req *http.Request) {
	if req.Method == "PUT" {
		var ctx = context.Background()
		code := req.URL.Query().Get("code")
		if code == "" {
			util.Response(w, nil, errors.New("promo code not define"))
			return
		}
		result, err := h.promoHandler.DecrimentPromoByCode(ctx, code)

		util.Response(w, result, err)
	} else {
		util.Response(w, nil, errors.New("error method"))
	}
}
