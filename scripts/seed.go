package scripts

import (
	"log"
	"siul-pbj-api/config"
	"siul-pbj-api/models"
)

func main() {
	config.ConnectDB()

	adminUser := models.User{
		Nama:         "Super Admin PEP",
		Username:     "admin_pep",
		PasswordHash: "adminpep",
		Role:         "pep",
	}

	// Insert user jika username admin_pep belum ada
	var existingUser models.User
	result := config.DB.Where("username = ?", "admin_pep").First(&existingUser)

	if result.Error != nil { // jika belum ada (record not found)
		config.DB.Create(&adminUser)
		log.Println("Berhasil membuat user admin_pep otomatis dengan password: adminpep")
	} else {
		// Update password sekalian memastikan sesuai "adminpep"
		existingUser.PasswordHash = "adminpep"
		config.DB.Save(&existingUser)
		log.Println("User admin_pep sudah ada, password telah di-reset ke: adminpep")
	}
}
