package model

type DiskonRequest struct {
	HargaAwal   float64 `json:"harga_awal"`
	PakeVoucher bool    `json:"pake_voucher"`
}

type DiskonResponse struct {
	HargaAwal          float64 `json:"harga_awal"`
	PakeVoucher        bool    `json:"pake_voucher"`
	HargaSetelahDiskon float64 `json:"harga_setelah_diskon"`
	PointsDidapatkan   float64 `json:"points_didapatkan"`
}
