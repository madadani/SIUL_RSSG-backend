package main

import (
	"log"

	"siul-pbj-api/internal/domain"
	"siul-pbj-api/pkg/config"
)

func main() {
	config.ConnectDB()

	seedUsers()
	seedMasterData()

	log.Println("✅ All seed data completed successfully!")
}

func seedUsers() {
	log.Println("🌱 Seeding users...")
	config.DB.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE")

	users := []domain.User{
		{Nama: "Super Admin PEP", Username: "admin_pep", PasswordHash: "adminpep", Role: "pep"},

		// PPKOM
		{Nama: "Bekti Nugroho, S.Kom., M.Eng.", Username: "bekti", PasswordHash: "ppkom123", Role: "ppkom"},
		{Nama: "Suramin, S.KM., S.Kep., Ns.", Username: "suramin", PasswordHash: "ppkom123", Role: "ppkom"},
		{Nama: "dr. Nofi Kusumaningrum", Username: "nofi", PasswordHash: "ppkom123", Role: "ppkom"},
		{Nama: "Lilik Subagyo, S.Kep., Ns.", Username: "lilik", PasswordHash: "ppkom123", Role: "ppkom"},
		{Nama: "dr. Mayasari Ayu Hendrawati", Username: "mayasari", PasswordHash: "ppkom123", Role: "ppkom"},
		{Nama: "Eko Sujianto, S.Kep., Ns.", Username: "eko", PasswordHash: "ppkom123", Role: "ppkom"},
		{Nama: "Dr. dr. Kinik Darsono, M.Pd.Ked.", Username: "kinik", PasswordHash: "ppkom123", Role: "ppkom"},
		{Nama: "Imas Wulandari, S.Kom., M.Eng.", Username: "imas", PasswordHash: "ppkom123", Role: "ppkom"},

		// PPTK
		{Nama: "Yuwono Laksito, A.Md.", Username: "yuwono", PasswordHash: "pptk123", Role: "pptk"},
		{Nama: "Andreas Hasta Riawan", Username: "andreas", PasswordHash: "pptk123", Role: "pptk"},
		{Nama: "Nunung Agus Dwi H., M.Si.", Username: "nunung", PasswordHash: "pptk123", Role: "pptk"},

		// PP
		{Nama: "Fitria Purnamawati, S.Kep., Ns., M.Kep.", Username: "fitria", PasswordHash: "pp123", Role: "pp"},
		{Nama: "Agus Dwiyanto, S.Farm., Apt.", Username: "agus", PasswordHash: "pp123", Role: "pp"},
		{Nama: "Purwanto, A.Md.", Username: "purwanto", PasswordHash: "pp123", Role: "pp"},
		{Nama: "Bagus Prasetya Utama, A.Md.TE.", Username: "bagus", PasswordHash: "pp123", Role: "pp"},
		{Nama: "Candra Nugraheni, S.ST.", Username: "candra-nugraheni", PasswordHash: "pp123", Role: "pp"},
		{Nama: "Candra Rosanti, S.KM.", Username: "candra-rosanti", PasswordHash: "pp123", Role: "pp"},
		{Nama: "Tulus Wahyuno, A.Md.", Username: "tulus", PasswordHash: "pp123", Role: "pp"},
	}

	for _, user := range users {
		var existing domain.User
		result := config.DB.Where("username = ?", user.Username).First(&existing)
		if result.Error != nil {
			config.DB.Create(&user)
			log.Printf("  ✓ Created user %s (%s)\n", user.Username, user.Role)
		} else {
			existing.Nama = user.Nama
			existing.PasswordHash = user.PasswordHash
			existing.Role = user.Role
			config.DB.Save(&existing)
			log.Printf("  ↻ Updated user %s\n", user.Username)
		}
	}
}

func seedMasterData() {
	log.Println("🌱 Seeding master categories & items...")

	categoriesData := map[string][]string{
		"BELANJA BARANG DAN JASA": {
			"Belanja Alat Tulis Kantor",
			"Belanja Alat Listrik dan Elektronik",
			"Belanja Perangko, Materai, dan Benda Pos",
			"Belanja Alat Kebersihan dan Bahan Pembersih",
			"Belanja Alat Rumah Tangga Kantor",
			"Belanja Hadiah Lomba / Penghargaan / Suvenir",
			"Belanja Perkakas Kerja",
			"Belanja Barang Pakai Habis Kesehatan",
			"Belanja Bahan Obat-Obatan",
			"Belanja Bahan Kimia/ Radiologi/ Oksigen/ PMI",
			"Belanja Bahan Laboratorium",
			"Belanja Cetak",
			"Belanja Penggandaan",
			"Belanja Makanan dan Minuman Harian Pegawai",
			"Belanja Makanan dan Minuman Rapat",
			"Belanja Makanan dan Minuman Tamu",
			"Belanja Makanan dan Minuman Pasien",
			"Belanja Pengisian Tabung Pemadam Kebakaran",
			"Belanja Pengisian Tabung Gas",
			"Belanja Telepon / Faksimil/ Internet/paket data",
			"Belanja Air",
			"Belanja Listrik",
			"Belanja Pakaian Dinas Harian",
			"Belanja Surat Kabar/Majalah",
			"Belanja Paket/Pengiriman",
			"Belanja Perjalanan Dinas Dalam Daerah",
			"Belanja Perjalanan Dinas Luar Daerah",
			"Belanja Penggantian Suku Cadang",
			"Belanja Bahan Bakar Minyak dan Pelumas",
			"Belanja Asuransi Kesehatan",
			"Belanja iuran jaminan kesehatan ASN",
			"Belanja Asuransi Kendaraan",
			"Belanja Sertifikasi / Kalibrasi",
			"Belanja Pajak/Retribusi",
			"Belanja Jasa Tenaga Kesehatan / Non ASN",
			"Belanja Jasa Tenaga Pengamanan dan Sopir",
			"Belanja Jasa Tenaga Administrasi",
			"Belanja Jasa Kebersihan",
			"Belanja Penguburan Jenazah dan Perlengkapannya",
			"Belanja Jasa Sosial",
			"Belanja Jasa Pengambilan/Pengumpulan/Pengangkutan Sampah",
			"Belanja Sewa",
			"Belanja Jasa Asuransi, Perbankan, dan Keuangan",
			"Belanja Jasa Penerangan, Iklan/Reklame,Film dan Pemotretan",
			"Belanja Jasa Konsultansi Manajemen/Keuangan/SDM",
			"Belanja Jasa Pelayanan Kesehatan bagi ASN",
			"Belanja Insentif bagi Pegawai Non ASN atas Pemungutan Retribusi Jasa Umum- Pelayanan Kesehatan",
			"Belanja Jasa Servis",
			"Belanja Jasa Pest Control",
			"Belanja Bimbingan Teknis, Sosialisasi, Pelatihan",
			"Belanja Pemeliharaan Alat Kantor dan Rumah Tangga lainnya",
			"Belanja Pemeliharaan Alat Pendingin Ruangan ( Air Conditioner )",
			"Belanja Pemeliharaan Alat Kedokteran dan Kesehatan umum",
			"Belanja Pemeliharaan Bangunan Gedung - Pemeliharaan Lift",
			"Belanja Pemeliharaan Bangunan Gedung Tempat Kerja",
			"Belanja Pemeliharaan Bangunan Tempat Kerja - Taman",
			"Belanja Pemeliharaan Alat Kantor dan Rumah Tangga - Elektronik dan Software",
			"Belanja Lainnya",
		},
		"BELANJA MODAL PERALATAN DAN MESIN": {
			"Belanja Modal Kendaraan Dinas",
			"Belanja Modal Alat Kantor Lainnya",
			"Belanja Modal Meubelair",
			"Belanja Modal Alat Dapur",
			"Belanja Modal Alat Rumah Tangga Lainnya",
			"Belanja Modal Alat Kesehatan Umum Lainnya",
			"Belanja Modal Peralatan Komputer",
			"Belanja Modal Alat Komunikasi",
		},
		"BELANJA MODAL GEDUNG DAN BANGUNAN": {
			"Belanja Modal Rehab Gedung Transit Dokter - Gedung eks Kelurahan",
			"Belanja Modal Ruang Intalasi Bedah Sentral",
		},
	}

	for catName, items := range categoriesData {
		var kode string
		switch catName {
		case "BELANJA BARANG DAN JASA":
			kode = "BBJ"
		case "BELANJA MODAL PERALATAN DAN MESIN":
			kode = "BMPM"
		case "BELANJA MODAL GEDUNG DAN BANGUNAN":
			kode = "BMGB"
		default:
			kode = "KAT-" + catName[:3]
		}

		var kategori domain.KategoriBelanja
		result := config.DB.Where("kode_kategori = ?", kode).First(&kategori)
		if result.Error != nil {
			kategori = domain.KategoriBelanja{KodeKategori: kode, NamaKategori: catName}
			config.DB.Create(&kategori)
			log.Printf("  ✓ Created Kategori: %s\n", catName)
		} else {
			kategori.NamaKategori = catName
			config.DB.Save(&kategori)
		}

		for _, itemName := range items {
			var masterBarang domain.MasterRincianBelanja
			res := config.DB.Where("nama_barang = ? AND kategori_belanja_id = ?", itemName, kategori.ID).First(&masterBarang)
			if res.Error != nil {
				masterBarang = domain.MasterRincianBelanja{
					KategoriBelanjaID: kategori.ID,
					NamaBarang:        itemName,
					Spesifikasi:       "-",
					Satuan:            "Paket/Kegiatan",
					HargaReferensi:    0.0,
				}
				config.DB.Create(&masterBarang)
				log.Printf("    → Created Barang: %s\n", itemName)
			}
		}
	}
}
