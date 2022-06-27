package models

type Peminjaman struct {
	Id              int    `gorm:"primaryKey;autoIncrement;" json:"id"`
	Nama            string `json:"nama"`
	No_telp         string `json:"no_telp"`
	Id_Buku         string `json:"id_buku"`
	No_Rak          string `json:"no_rak"`
	Tanggal_Kembali string `json:"tgl_kembali"`
}
