package controllers

import (
	"net/http"
	"perpustakaan3/Transaksi/usecase"
)

type ctrl struct {
	us usecase.Usecase
}

type ControllerInterface interface {
	GetDataPeminjaman(w http.ResponseWriter, r *http.Request)
	PostDataPeminjaman(w http.ResponseWriter, r *http.Request)
	GetDetailDataPeminjaman(w http.ResponseWriter, r *http.Request)
	PutDataPeminjaman(w http.ResponseWriter, r *http.Request)
	DeleteDataPeminjamann(w http.ResponseWriter, r *http.Request)

	GetDataPengembalian(w http.ResponseWriter, r *http.Request)
	PostDataPengembalian(w http.ResponseWriter, r *http.Request)
	GetDetailPengembalian(w http.ResponseWriter, r *http.Request)
	PutDataPengembalian(w http.ResponseWriter, r *http.Request)
	DeleteDataPengembalian(w http.ResponseWriter, r *http.Request)
}

func NewController(us usecase.Usecase) ControllerInterface {
	return &ctrl{us: us}
}
