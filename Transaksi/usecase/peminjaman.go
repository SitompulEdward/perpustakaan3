package usecase

import (
	"errors"
	"perpustakaan3/Transaksi/models"
)

func (r *UC) GetAllPeminjaman() ([]models.Peminjaman, error) {
	var modelpmj []models.Peminjaman

	err := r.queryrepo.SelectAllPeminjaman(modelpmj)
	if err != nil {
		return modelpmj, err
	}

	return modelpmj, nil
}

func (r *UC) InsertDataPeminjaman(data models.Peminjaman) error {
	err := r.queryrepo.InsertDataPeminjaman(&data)
	if err != nil {
		return err
	}
	return nil
}

func (r *UC) SelectDetailPeminjaman(id int) (data models.Peminjaman, err error) {
	data, err = r.queryrepo.SelectPeminjamanById(id)
	if err != nil {
		return data, errors.New("data tidak ada")
	}

	return data, nil
}

func (r *UC) UpdateDataPeminjaman(data models.Peminjaman) error {
	pmj, err := r.queryrepo.SelectPeminjamanById(data.Id)
	if err != nil {
		return errors.New("admin tidak ditemukan")
	}

	pmj.Nama = data.Nama
	pmj.No_telp = data.No_telp
	pmj.Id_Buku = data.Id_Buku
	pmj.No_Rak = data.No_Rak
	pmj.Tanggal_Kembali = data.Tanggal_Kembali

	err = r.queryrepo.UpdatePeminjaman(pmj)
	if err != nil {
		return err
	}

	return nil
}

func (r *UC) DeleteDataPeminjaman(data models.Peminjaman) error {
	pmj, err := r.queryrepo.SelectPeminjamanById(data.Id)
	if err != nil {
		return errors.New("admin tidak ditemukan")
	}

	pmj.Id = data.Id

	err = r.queryrepo.DeletePeminjaman(pmj)
	if err != nil {
		return err
	}

	return nil
}
