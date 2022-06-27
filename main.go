package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"perpustakaan3/Transaksi/connection"
	"perpustakaan3/Transaksi/controllers"
	"perpustakaan3/Transaksi/repositories"
	"perpustakaan3/Transaksi/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error Load Config File !")
	}
}

func main() {
	koneksidb := connection.ConnectToDB()
	repo := repositories.NewRepo(koneksidb)
	usecase := usecase.NewUsecase(repo)
	ctrl := controllers.NewController(usecase)

	r := chi.NewRouter()

	r.Get("/get-data-peminjaman", ctrl.GetDataPeminjaman)
	r.Post("/post-data-peminjaman", ctrl.PostDataPeminjaman)
	r.Get("/get-data-detail-peminjaman/{id}", ctrl.GetDetailDataPeminjaman)
	r.Put("/update-data-peminjaman/{id}", ctrl.PutDataPeminjaman)
	r.Delete("/delete-data-peminjaman/{id}", ctrl.DeleteDataPeminjamann)

	r.Get("/get-data-pengembalian", ctrl.GetDataPengembalian)
	r.Post("/post-data-pengembalian", ctrl.PostDataPengembalian)
	r.Get("/get-detail-data-pengembalian/{id}", ctrl.GetDetailPengembalian)
	r.Put("/update-data-pengembalian/{id}", ctrl.PutDataPengembalian)
	r.Delete("/delete-data-pengembalian/{id}", ctrl.DeleteDataPengembalian)

	if err := http.ListenAndServe(":"+os.Getenv("HOST")+"", r); err != nil {
		fmt.Println("Error Starting Service !")
	}
}
