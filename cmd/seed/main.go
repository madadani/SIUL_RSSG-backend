package main

import (
	"fmt"
	"log"
	"os"

	"siul-pbj-api/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	_ = godotenv.Load()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi DB:", err)
	}

	log.Println("Connected to DB for seeding...")

	// ==================== SEED USERS ====================
	users := []models.User{
		{Nama: "Super Admin PEP", Username: "admin_pep", PasswordHash: "adminpep", Role: "pep"},
		{Nama: "Bapak Agus PPTK", Username: "admin_pptk", PasswordHash: "adminpptk", Role: "pptk"},
		{Nama: "Ibu Sari PPKOM", Username: "admin_ppkom", PasswordHash: "adminppkom", Role: "ppkom"},
		{Nama: "Bapak Rudi PP", Username: "admin_pp", PasswordHash: "adminpp", Role: "pp"},
	}

	for _, u := range users {
		var existing models.User
		result := db.Where("username = ?", u.Username).First(&existing)
		if result.Error != nil {
			// User does not exist, create it
			db.Create(&u)
			log.Printf("✅ Created user: %s (%s)\n", u.Nama, u.Role)
		} else {
			log.Printf("⏩ User sudah ada: %s (%s)\n", existing.Nama, existing.Role)
		}
	}

	// ==================== SEED KATEGORI BELANJA ====================
	kategoris := []models.KategoriBelanja{
		{KodeKategori: "ATK", NamaKategori: "Alat Tulis Kantor"},
		{KodeKategori: "IT", NamaKategori: "Peralatan IT & Komputer"},
		{KodeKategori: "JASA", NamaKategori: "Jasa Konsultansi"},
	}

	for _, k := range kategoris {
		var existing models.KategoriBelanja
		result := db.Where("kode_kategori = ?", k.KodeKategori).First(&existing)
		if result.Error != nil {
			db.Create(&k)
			log.Printf("✅ Created kategori: %s\n", k.NamaKategori)
		} else {
			log.Printf("⏩ Kategori sudah ada: %s\n", existing.NamaKategori)
		}
	}

	log.Println("\n🎉 Seeding selesai!")
	log.Println("================================")
	log.Println("AKUN LOGIN YANG TERSEDIA:")
	log.Println("================================")
	log.Println("PEP      : admin_pep    / adminpep")
	log.Println("PPTK     : admin_pptk   / adminpptk")
	log.Println("PPKOM    : admin_ppkom  / adminppkom")
	log.Println("PP       : admin_pp     / adminpp")
	log.Println("================================")
}
