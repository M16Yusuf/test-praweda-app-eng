package handler

import (
	"net/http"
	"test-praweda-app-eng/internal/model"

	"github.com/gin-gonic/gin"
)

type EcommerceHandler struct {
}

func NewEcommerceHandler() *EcommerceHandler {
	return &EcommerceHandler{}
}

func (h *EcommerceHandler) DiskonCalculation(ctx *gin.Context) {
	var body model.DiskonRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			IsSuccess: false,
			Code:      http.StatusBadRequest,
			Err:       "invalid input",
		})
		return
	}

	if !body.PakeVoucher {
		ctx.JSON(http.StatusOK, model.Response{
			IsSuccess: true,
			Code:      http.StatusOK,
			Msg:       "anda tidak mendapatkan point karena tidak menggunakan voucher",
		})
		return
	}

	if body.PakeVoucher {
		hargaSetelahDiskon := body.HargaAwal * 0.5
		pointsDidapatkan := hargaSetelahDiskon * 0.02

		response := model.DiskonResponse{
			HargaAwal:          body.HargaAwal,
			PakeVoucher:        body.PakeVoucher,
			HargaSetelahDiskon: hargaSetelahDiskon,
			PointsDidapatkan:   pointsDidapatkan,
		}
		ctx.JSON(http.StatusOK, model.ResponseData{
			Response: model.Response{
				IsSuccess: true,
				Code:      http.StatusOK,
				Msg:       "Anda menggunakan voucher diskon 50% dan mendapatkan points sebesar 2% dari harga setelah diskon",
			},
			Data: response,
		})
	}
}
