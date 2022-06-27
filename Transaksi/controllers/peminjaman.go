package controllers

import (
	"encoding/json"
	"net/http"
	"perpustakaan3/Transaksi/connection"
	"perpustakaan3/Transaksi/models"
	"strconv"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB = connection.ConnectToDB()
}

func (c *ctrl) GetDataPeminjaman(w http.ResponseWriter, r *http.Request) {
	var listdatapmj []models.Peminjaman

	DB.Find(&listdatapmj)
	datajson, err := json.Marshal(listdatapmj)

	if err != nil {
		w.Write([]byte("Error Convert TO JSON"))
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(datajson)
	w.WriteHeader(200)
	return
}

func (c *ctrl) PostDataPeminjaman(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var datarequest models.Peminjaman

	err := decoder.Decode(&datarequest)

	if err != nil {
		w.Write([]byte("Error Decode JSON Payload"))
		w.WriteHeader(500)
		return
	}

	err = c.us.InsertDataPeminjaman(datarequest)
	if err != nil {
		ResponseApi(w, 500, nil, "Internal Server Error")
		return
	}

	ResponseApi(w, 200, nil, "Sukses Insert Data")
	return
}

func (c *ctrl) GetDetailDataPeminjaman(w http.ResponseWriter, r *http.Request) {
	idpmj := chi.URLParam(r, "id")

	idConvert, err := strconv.Atoi(idpmj)

	if err != nil {
		ResponseApi(w, 500, nil, "Internal Server Error")
		return
	}

	data, err := c.us.SelectDetailPeminjaman(idConvert)
	if err != nil {
		ResponseApi(w, 500, nil, err.Error())
	}

	ResponseApi(w, 200, data, "Sukses Get Data")
}

func (c *ctrl) PutDataPeminjaman(w http.ResponseWriter, r *http.Request) {
	idadmin := chi.URLParam(r, "id")

	decoder := json.NewDecoder(r.Body)
	var datarequest models.Peminjaman
	err := decoder.Decode(&datarequest)

	if err != nil {
		ResponseApi(w, 500, nil, "INTERNAL SERVER ERROR")
		return
	}

	idconvert, err := strconv.Atoi(idadmin)
	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}

	datarequest.Id = idconvert
	err = c.us.UpdateDataPeminjaman(datarequest)
	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}

	ResponseApi(w, 200, nil, "Sukses Update Data")
	return
}

func (c *ctrl) DeleteDataPeminjamann(w http.ResponseWriter, r *http.Request) {
	idpmj := chi.URLParam(r, "id")

	if idpmj == "" {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}

	var datarequest models.Peminjaman
	idConvert, err := strconv.Atoi(idpmj)

	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}

	datarequest.Id = idConvert
	err = c.us.DeleteDataPeminjaman(datarequest)

	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}

	ResponseApi(w, 200, nil, "Sukses delete")
	return
}

func ResponseApi(w http.ResponseWriter, code int, data interface{}, msg string) {

	resevice := models.Response{}
	resevice.Code = code
	resevice.Message = msg
	resevice.Data = data

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resevice)

}
