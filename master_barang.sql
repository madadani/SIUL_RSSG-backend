-- phpMyAdmin SQL Dump
-- version 5.2.2
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Mar 29, 2026 at 08:24 AM
-- Server version: 8.0.30
-- PHP Version: 7.4.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `detail_anggaran.sql`
--

-- --------------------------------------------------------

--
-- Table structure for table `master_barang`
--

CREATE TABLE `master_barang` (
  `id` bigint UNSIGNED NOT NULL,
  `nama` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `kategori` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `master_barang`
--

INSERT INTO `master_barang` (`id`, `nama`, `kategori`, `created_at`, `updated_at`) VALUES
(1, 'Laptop', 'Belanja Modal Peralatan Komputer', '2026-03-28 22:19:15', '2026-03-28 22:19:15'),
(2, 'Printer', 'Belanja Modal Peralatan Komputer', '2026-03-28 22:19:15', '2026-03-28 22:19:15'),
(3, 'Printer Barcode (XP360B)', 'Belanja Modal Peralatan Komputer', '2026-03-28 22:19:15', '2026-03-28 22:19:15'),
(4, 'Printer Epson L 121', 'Belanja Modal Peralatan Komputer', '2026-03-28 22:19:15', '2026-03-28 22:19:15'),
(5, 'Printer Passbook', 'Belanja Modal Peralatan Komputer', '2026-03-28 22:19:15', '2026-03-28 22:19:15'),
(6, 'Server DELL PowerEdge R760xs', 'Belanja Modal Peralatan Komputer', '2026-03-28 22:19:15', '2026-03-28 22:19:15'),
(7, 'Paket Peralatan Komputer (CPU, Printer, NAS, Hardisk)', 'Belanja Modal Peralatan Komputer', '2026-03-28 22:19:15', '2026-03-28 22:19:15'),
(8, 'Printer Thermal IWARE IW-58AC 58mm', 'Belanja Modal Peralatan Komputer', '2026-03-28 22:19:15', '2026-03-28 22:19:15'),
(9, 'HP Infinix Smart 9', 'Belanja Modal Alat Komunikasi', '2026-03-28 22:19:15', '2026-03-28 22:19:15'),
(10, 'Smartphone Infinix Hot 60i (6GB/128GB)', 'Belanja Modal Alat Komunikasi', '2026-03-28 22:19:15', '2026-03-28 22:19:15'),
(11, 'Smartphone Infinix 10 (4GB/128GB) - TPPRI', 'Belanja Modal Alat Komunikasi', '2026-03-28 22:19:15', '2026-03-28 22:19:15'),
(12, 'Smartphone Infinix 10 (4GB/128GB) - Bangsal Vanda', 'Belanja Modal Alat Komunikasi', '2026-03-28 22:19:15', '2026-03-28 22:19:15'),
(13, 'Smartphone Infinix 10 Plus (8GB/128GB) - Instalasi Radiologi', 'Belanja Modal Alat Komunikasi', '2026-03-28 22:19:15', '2026-03-28 22:19:15'),
(14, 'Smartphone Infinix 10 (4GB/128GB) - Layanan Homecare', 'Belanja Modal Alat Komunikasi', '2026-03-28 22:19:15', '2026-03-28 22:19:15'),
(15, 'Vein Viewer', 'Belanja Modal Alat Kesehatan Umum Lainnya', '2026-03-28 22:19:15', '2026-03-28 22:19:15'),
(16, 'Ambubag', 'Belanja Modal Alat Kesehatan Umum Lainnya', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(17, 'Stetoscope', 'Belanja Modal Alat Kesehatan Umum Lainnya', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(18, 'Nebulizer', 'Belanja Modal Alat Kesehatan Umum Lainnya', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(19, 'Oximeter', 'Belanja Modal Alat Kesehatan Umum Lainnya', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(20, 'Termometer', 'Belanja Modal Alat Kesehatan Umum Lainnya', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(21, 'Infus Pump', 'Belanja Modal Alat Kesehatan Umum Lainnya', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(22, 'Syring Pump', 'Belanja Modal Alat Kesehatan Umum Lainnya', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(23, 'Pasien Monitor', 'Belanja Modal Alat Kesehatan Umum Lainnya', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(24, 'Pengawaan Simcoe, Sinskey, Colibri', 'Belanja Modal Alat Kesehatan Umum Lainnya', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(25, 'Regulator Oksigen Central', 'Belanja Modal Alat Kesehatan Umum Lainnya', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(26, 'Troli Instrument', 'Belanja Modal Alat Kesehatan Umum Lainnya', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(27, 'Tiang Infus', 'Belanja Modal Alat Kesehatan Umum Lainnya', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(28, 'Meja Makan Mayo', 'Belanja Modal Alat Kesehatan Umum Lainnya', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(29, 'Bed Side Cabinet', 'Belanja Modal Alat Kesehatan Umum Lainnya', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(30, 'Kursi Tunggu', 'Belanja Modal Meubelair', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(31, 'Sofabed', 'Belanja Modal Meubelair', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(32, 'Kursi Putar', 'Belanja Modal Meubelair', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(33, 'Meja Kerja', 'Belanja Modal Meubelair', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(34, 'Barang Meubelair (umum)', 'Belanja Modal Meubelair', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(35, 'LED TV 32 inch (2 Unit)', 'Belanja Modal Rumah Tangga Lainnya', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(36, 'LED TV 43 inch (3 Unit)', 'Belanja Modal Rumah Tangga Lainnya', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(37, 'Dispenser Galon Bawah (1 Unit)', 'Belanja Modal Rumah Tangga Lainnya', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(38, 'Kulkas 1 Pintu (3 Unit)', 'Belanja Modal Rumah Tangga Lainnya', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(39, 'Trolly Stainless (2 pcs)', 'Belanja Modal Rumah Tangga Lainnya', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(40, 'Gorden Ruang Vanda dan PICU Anggrek', 'Belanja Modal Rumah Tangga Lainnya', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(41, 'Gorden Unit Hemodialisis/HD', 'Belanja Modal Rumah Tangga Lainnya', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(42, 'Mesin Pemotong Rumput HONDA UMR435N', 'Belanja Modal Rumah Tangga Lainnya', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(43, 'Pallet Plastik 1200 x 1000x 150 mm (6 Unit)', 'Belanja Modal Rumah Tangga Lainnya', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(44, 'Showcase Polytron SCN-1020', 'Belanja Modal Rumah Tangga Lainnya', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(45, 'Rice Cooker Miyako MCG-171 6 Lt', 'Belanja Modal Rumah Tangga Lainnya', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(46, 'Tempat Sampah Injak', 'Belanja Modal Rumah Tangga Lainnya', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(47, 'AC Spilt Panasonic 3/4 PK (Geriatri)', 'Belanja Modal Alat Kantor Lainnya', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(48, 'Air Conditioner (Unit Dialisis, ICU, dan IBS)', 'Belanja Modal Alat Kantor Lainnya', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(49, 'Platform Hand Truck 300Kg Plastic Black', 'Belanja Modal Alat Kantor Lainnya', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(50, 'Cup Sealer', 'Belanja Modal Alat Dapur', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(51, 'Troli Pengantar Makanan', 'Belanja Modal Alat Dapur', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(52, 'Meja Stainless', 'Belanja Modal Alat Dapur', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(53, 'Pembangunan Kantin RSUD', 'Belanja Modal Gedung dan Bangunan BLUD', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(54, 'Rehab Gedung Limbah', 'Belanja Modal Gedung dan Bangunan BLUD', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(55, 'Rehab Atap dan Selasar Ruang Hemodialisa', 'Belanja Modal Gedung dan Bangunan BLUD', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(56, 'Instalasi Pipa Gas Unit Dialisis', 'Belanja Modal Instalasi Lain', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(57, 'Sapu Lidi', 'Belanja Alat Rumah Tangga Kantor', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(58, 'Bendera Merah Putih 180x120', 'Belanja Alat Rumah Tangga Kantor', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(59, 'Plastik Sampah', 'Belanja Alat Rumah Tangga Kantor', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(60, 'Plastik Klip', 'Belanja Alat Rumah Tangga Kantor', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(61, 'Gelas Plastik', 'Belanja Alat Rumah Tangga Kantor', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(62, 'Sendok Plastik', 'Belanja Alat Rumah Tangga Kantor', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(63, 'Matras Puzzle Evamat', 'Belanja Alat Rumah Tangga Kantor', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(64, 'Mainan Kuda-kudaan Anak', 'Belanja Alat Rumah Tangga Kantor', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(65, 'Mainan Sepeda Anak', 'Belanja Alat Rumah Tangga Kantor', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(66, 'Stiker Kaca', 'Belanja Alat Rumah Tangga Kantor', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(67, 'Kontainer Bok', 'Belanja Alat Rumah Tangga Kantor', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(68, 'Kursi Tunggu Besi', 'Belanja Alat Rumah Tangga Kantor', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(69, 'Plastik Kebutuhan Instalasi Gizi', 'Belanja Alat Rumah Tangga Kantor', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(70, 'Sabun Mandi Bayi', 'Belanja Alat Rumah Tangga Kantor', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(71, 'Properti Lomba Pitoelasan RSSG', 'Belanja Alat Rumah Tangga Kantor', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(72, 'Perkakas Makan Instalasi Gizi', 'Belanja Alat Rumah Tangga Kantor', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(73, 'Manekin Cowok', 'Belanja Alat Rumah Tangga Kantor', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(74, 'Properti Suporter Lomba 17-an', 'Belanja Alat Rumah Tangga Kantor', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(75, 'Lid Cup Sealer', 'Belanja Alat Rumah Tangga Kantor', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(76, 'Gagang Stampel', 'Belanja Alat Rumah Tangga Kantor', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(77, 'Ember', 'Belanja Alat Rumah Tangga Kantor', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(78, 'LED Downlight', 'Belanja Alat Listrik dan Elektronik', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(79, 'Baterai AA', 'Belanja Alat Listrik dan Elektronik', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(80, 'STB Lubby', 'Belanja Alat Listrik dan Elektronik', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(81, 'Lampu LED', 'Belanja Alat Listrik dan Elektronik', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(82, 'Baterai Alkaline AA, AAA', 'Belanja Alat Listrik dan Elektronik', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(83, 'Tinta Refill Catridge 85A dan 12A', 'Belanja Alat Tulis Kantor (ATK)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(84, 'Kertas HVS F4 60 gr', 'Belanja Alat Tulis Kantor (ATK)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(85, 'Amplop Kecil', 'Belanja Alat Tulis Kantor (ATK)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(86, 'Kertas HVS F4', 'Belanja Alat Tulis Kantor (ATK)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(87, 'Label Thermall', 'Belanja Alat Tulis Kantor (ATK)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(88, 'Buku Instrumen Akreditasi RS', 'Belanja Alat Tulis Kantor (ATK)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(89, 'Kertas HVS', 'Belanja Alat Tulis Kantor (ATK)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(90, 'Sabun Cuci Piring', 'Belanja Alat Kebersihan dan Bahan Pembersih', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(91, 'Safety Box', 'Belanja Alat Kebersihan dan Bahan Pembersih', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(92, 'Desinfektan', 'Belanja Alat Kebersihan dan Bahan Pembersih', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(93, 'Detergen', 'Belanja Alat Kebersihan dan Bahan Pembersih', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(94, 'Baymed Plus', 'Belanja Alat Kebersihan dan Bahan Pembersih', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(95, 'Perhydrol Forte White Hydrogen', 'Belanja Alat Kebersihan dan Bahan Pembersih', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(96, 'Acid Citric Monohydrate', 'Belanja Alat Kebersihan dan Bahan Pembersih', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(97, 'Pouches', 'Belanja Alat Kebersihan dan Bahan Pembersih', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(98, 'Feeding Tube', 'Belanja Barang Pakai Habis Kesehatan (BPHK) (Contoh Barang)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(99, 'Guedel Airway', 'Belanja Barang Pakai Habis Kesehatan (BPHK) (Contoh Barang)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(100, 'Kasa Lipat', 'Belanja Barang Pakai Habis Kesehatan (BPHK) (Contoh Barang)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(101, 'Kapas', 'Belanja Barang Pakai Habis Kesehatan (BPHK) (Contoh Barang)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(102, 'Apron Sleeveless', 'Belanja Barang Pakai Habis Kesehatan (BPHK) (Contoh Barang)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(103, 'Masker', 'Belanja Barang Pakai Habis Kesehatan (BPHK) (Contoh Barang)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(104, 'Sarung Tangan (Latex Exam Gloves, Surgical Glove)', 'Belanja Barang Pakai Habis Kesehatan (BPHK) (Contoh Barang)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(105, 'Spuit/Syringe', 'Belanja Barang Pakai Habis Kesehatan (BPHK) (Contoh Barang)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(106, 'Infusion Set', 'Belanja Barang Pakai Habis Kesehatan (BPHK) (Contoh Barang)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(107, 'IV Cannula', 'Belanja Barang Pakai Habis Kesehatan (BPHK) (Contoh Barang)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(108, 'Oksigen Mask / Nasal Cannula', 'Belanja Barang Pakai Habis Kesehatan (BPHK) (Contoh Barang)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(109, 'Kasa Steril', 'Belanja Barang Pakai Habis Kesehatan (BPHK) (Contoh Barang)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(110, 'Plester', 'Belanja Barang Pakai Habis Kesehatan (BPHK) (Contoh Barang)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(111, 'Blood Lancet', 'Belanja Barang Pakai Habis Kesehatan (BPHK) (Contoh Barang)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(112, 'ECG Electrode', 'Belanja Barang Pakai Habis Kesehatan (BPHK) (Contoh Barang)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(113, 'Benang Jahit Bedah', 'Belanja Barang Pakai Habis Kesehatan (BPHK) (Contoh Barang)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(114, 'Balloon Catheter', 'Belanja Barang Pakai Habis Kesehatan (BPHK) (Contoh Barang)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(115, 'Kertas EKG', 'Belanja Barang Pakai Habis Kesehatan (BPHK) (Contoh Barang)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(116, 'Cairan Injeksi', 'Belanja Barang Pakai Habis Kesehatan (BPHK) (Contoh Barang)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(117, 'Underpad', 'Belanja Barang Pakai Habis Kesehatan (BPHK) (Contoh Barang)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(118, 'Alkohol Swab', 'Belanja Barang Pakai Habis Kesehatan (BPHK) (Contoh Barang)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(119, 'Urine Bag', 'Belanja Barang Pakai Habis Kesehatan (BPHK) (Contoh Barang)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(120, 'Obat generik dan paten (contoh: Lansoprazole, Simvastatin, Amoxicillin, Paracetamol, dll)', 'Belanja Bahan Obat-obatan (Contoh Barang)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(121, 'Infus (NaCl 0.9%, Dextrose, dll)', 'Belanja Bahan Obat-obatan (Contoh Barang)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(122, 'Vitamin', 'Belanja Bahan Obat-obatan (Contoh Barang)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(123, 'Antibiotik', 'Belanja Bahan Obat-obatan (Contoh Barang)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(124, 'Gas O2', 'Belanja Bahan Kimia/Radiologi/Oksigen/PMI', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(125, 'Liquid O2', 'Belanja Bahan Kimia/Radiologi/Oksigen/PMI', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(126, 'Darah dari PMI', 'Belanja Bahan Kimia/Radiologi/Oksigen/PMI', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(127, 'Bahan Radiologi (Film Drystar, Fuji Film DI-HL, Lopamiro, Barium Sulfat)', 'Belanja Bahan Kimia/Radiologi/Oksigen/PMI', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(128, 'Bahan Laboratorium (reagen)', 'Belanja Bahan Kimia/Radiologi/Oksigen/PMI', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(129, 'Bahan Laborat (umum untuk pemeriksaan)', 'Belanja Bahan Laboratorium', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(130, 'Pemeriksaan Hematologi Lengkap 5 Diff', 'Belanja Bahan Laboratorium', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(131, 'Pemeriksaan Elektrolit', 'Belanja Bahan Laboratorium', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(132, 'Air Minum (Karyawan & Pasien)', 'Belanja Makanan dan Minuman', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(133, 'Susu Entramix / Diabetasol', 'Belanja Makanan dan Minuman', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(134, 'Proten Rasa Vanila', 'Belanja Makanan dan Minuman', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(135, 'Bahan Makanan Pasien (Kena & Tidak Kena Pajak)', 'Belanja Makanan dan Minuman', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(136, 'Snack Dokter / Jaga Malam', 'Belanja Makanan dan Minuman', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(137, 'Konsumsi Rapat', 'Belanja Makanan dan Minuman', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(138, 'Konsumsi Tamu', 'Belanja Makanan dan Minuman', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(139, 'Cetak MMT, Spanduk, X-Banner, Backdrop', 'Belanja Cetak dan Penggandaan', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(140, 'Cetak Amplop, Stopmap, Berkas', 'Belanja Cetak dan Penggandaan', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(141, 'Cetak Akrilik (Petunjuk Arah, Zona Integritas)', 'Belanja Cetak dan Penggandaan', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(142, 'Cetak Stiker', 'Belanja Cetak dan Penggandaan', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(143, 'Sewa & Tagihan Fotocopy', 'Belanja Cetak dan Penggandaan', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(144, 'Jasa: Servis Kendaraan, Pest Control, Outsourcing (Tenaga Administrasi, Cleaning Service, Pengamanan, Sopir), Pengangkutan Sampah, Konsultasi.', 'Belanja Lainnya (Jasa, Suku Cadang, dll)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(145, 'Suku Cadang & Pemeliharaan: Sparepart Komputer, AC, Alat Kesehatan, RO, Pompa, Baut, Kabel, Aki.', 'Belanja Lainnya (Jasa, Suku Cadang, dll)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(146, 'Sewa: Sewa Kursi, Sewa Copy Machine.', 'Belanja Lainnya (Jasa, Suku Cadang, dll)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(147, 'Perjalanan Dinas: Biaya Perjalanan Dinas Dalam/Luar Daerah.', 'Belanja Lainnya (Jasa, Suku Cadang, dll)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(148, 'Pajak & Retribusi: Pajak Kendaraan, Pajak Air Tanah, Retribusi Sampah.', 'Belanja Lainnya (Jasa, Suku Cadang, dll)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(149, 'Komunikasi & Data: Paket Data, Pulsa, Tagihan Telpon/Internet, Hosting, Domain.', 'Belanja Lainnya (Jasa, Suku Cadang, dll)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(150, 'Insentif & Honor: Honorarium Tenaga Kesehatan, Non-ASN, Dewan Pengawas.', 'Belanja Lainnya (Jasa, Suku Cadang, dll)', '2026-03-28 22:19:16', '2026-03-28 22:19:16'),
(151, 'Sosial & Sumbangan: Dana Sosial, Sponsorship, Karangan Bunga, Santunan.', 'Belanja Lainnya (Jasa, Suku Cadang, dll)', '2026-03-28 22:19:16', '2026-03-28 22:19:16');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `master_barang`
--
ALTER TABLE `master_barang`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `master_barang`
--
ALTER TABLE `master_barang`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=152;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
