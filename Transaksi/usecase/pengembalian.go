package usecase

import (
	"errors"
	"perpustakaan3/Transaksi/models"
)

func (r *UC) GetAllPengembalian() ([]models.Pengembalian, error) {
	var modelpmj []models.Pengembalian

	err := r.queryrepo.SelectAllPengembalian(modelpmj)
	if err != nil {
		return modelpmj, err
	}

	return modelpmj, nil
}

func (r *UC) InsertDataPengembalian(data models.Pengembalian) error {
	err := r.queryrepo.InsertPengembalian(&data)
	if err != nil {
		return err
	}
	return nil
}

func (r *UC) SelectDetailPengembalian(id int) (data models.Pengembalian, err error) {
	data, err = r.queryrepo.SelectPengembalianById(id)
	if err != nil {
		return data, errors.New("data tidak ada")
	}

	return data, nil
}

func (r *UC) UpdateDataPengembalian(data models.Pengembalian) error {
	pmj, err := r.queryrepo.SelectPengembalianById(data.Id)
	if err != nil {
		return errors.New("data tidak ditemukan")
	}

	pmj.Nama = data.Nama
	pmj.No_telp = data.No_telp
	pmj.Id_Buku = data.Id_Buku
	pmj.No_Rak = data.No_Rak
	pmj.Tanggal_Kembali = data.Tanggal_Kembali

	err = r.queryrepo.UpdatePengembalian(pmj)
	if err != nil {
		return err
	}

	return nil
}

func (r *UC) DeleteDataPengembalian(data models.Pengembalian) error {
	pmj, err := r.queryrepo.SelectPengembalianById(data.Id)
	if err != nil {
		return errors.New("data tidak ditemukan")
	}

	pmj.Id = data.Id

	err = r.queryrepo.DeletePengembalian(pmj)
	if err != nil {
		return err
	}

	return nil
}
