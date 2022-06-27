package repositories

import "perpustakaan3/Transaksi/models"

func (r *repo) SelectAllPeminjaman(i interface{}) error {
	result := r.apps.Find(i)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repo) InsertDataPeminjaman(i interface{}) error {
	result := r.apps.Create(i)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repo) SelectPeminjamanById(id int) (data models.Peminjaman, err error) {
	err = r.apps.First(&data, id).Error
	return
}

func (r *repo) UpdatePeminjaman(i models.Peminjaman) error {
	result := r.apps.Save(&i)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repo) DeletePeminjaman(i models.Peminjaman) error {
	result := r.apps.Delete(i)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *repo) SelectAllPengembalian(i interface{}) error {
	result := r.apps.Find(i)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repo) InsertPengembalian(i interface{}) error {
	result := r.apps.Create(i)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repo) SelectPengembalianById(id int) (data models.Pengembalian, err error) {
	err = r.apps.First(&data, id).Error
	return
}

func (r *repo) UpdatePengembalian(i models.Pengembalian) error {
	result := r.apps.Save(&i)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *repo) DeletePengembalian(i models.Pengembalian) error {
	result := r.apps.Delete(i)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
