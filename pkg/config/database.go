package config

import (
	"fmt"
	"log"
	"os"

	"siul-pbj-api/internal/domain"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Note: Error loading .env file (might be using system env)")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database: ", err)
	}

	log.Println("Berhasil terhubung ke database PostgreSQL (siul_db) via GORM")

	// Auto Migrate (Update Schema)
	db.AutoMigrate(
		&domain.User{},
		&domain.KategoriBelanja{},
		&domain.MasterRincianBelanja{},
		&domain.DetailAnggaran{},
		&domain.Usulan{},
		&domain.RiwayatUsulan{},
		&domain.KewenanganPegawai{},
		&domain.RealisasiLaporan{},
	)

	DB = db
}
