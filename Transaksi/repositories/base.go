package repositories

import (
	"perpustakaan3/Transaksi/models"

	"gorm.io/gorm"
)

type repo struct {
	apps *gorm.DB
}

type Repo interface {
	SelectAllPeminjaman(i interface{}) error
	InsertDataPeminjaman(i interface{}) error
	SelectPeminjamanById(id int) (data models.Peminjaman, err error)
	UpdatePeminjaman(i models.Peminjaman) error
	DeletePeminjaman(i models.Peminjaman) error

	SelectAllPengembalian(i interface{}) error
	InsertPengembalian(i interface{}) error
	SelectPengembalianById(id int) (data models.Pengembalian, err error)
	UpdatePengembalian(i models.Pengembalian) error
	DeletePengembalian(i models.Pengembalian) error
}

func NewRepo(apps *gorm.DB) Repo {
	return &repo{apps}
}
