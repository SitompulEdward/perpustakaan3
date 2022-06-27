package usecase

import (
	"perpustakaan3/Transaksi/models"
	"perpustakaan3/Transaksi/repositories"
)

type UC struct {
	queryrepo repositories.Repo
}

type Usecase interface {
	GetAllPeminjaman() ([]models.Peminjaman, error)
	InsertDataPeminjaman(models.Peminjaman) error
	SelectDetailPeminjaman(id int) (data models.Peminjaman, err error)
	UpdateDataPeminjaman(models.Peminjaman) error
	DeleteDataPeminjaman(models.Peminjaman) error

	GetAllPengembalian() ([]models.Pengembalian, error)
	InsertDataPengembalian(models.Pengembalian) error
	SelectDetailPengembalian(id int) (data models.Pengembalian, err error)
	UpdateDataPengembalian(models.Pengembalian) error
	DeleteDataPengembalian(models.Pengembalian) error
}

func NewUsecase(r repositories.Repo) Usecase {
	return &UC{queryrepo: r}
}
