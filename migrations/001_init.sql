-- This file contains SQL snippets extracted from old fix/diagnose/reset Go scripts.
-- These snippets demonstrate how ENUMs and base tables were initialized for SIUL PBJ.
-- GORM AutoMigrate usually handles the schema, but ENUMs require manual SQL execution.

-- 1. DROP TAHAP UNTUK RESET KESELURUHAN (Optional, dari nuclear_fix.go)
/*
DROP TABLE IF EXISTS "realisasi_laporan" CASCADE;
DROP TABLE IF EXISTS "riwayat_usulan" CASCADE;
DROP TABLE IF EXISTS "kewenangan_pegawai" CASCADE;
DROP TABLE IF EXISTS "usulan" CASCADE;
DROP TABLE IF EXISTS "detail_anggaran" CASCADE;
DROP TABLE IF EXISTS "master_rincian_belanja" CASCADE;
DROP TABLE IF EXISTS "kategori_belanja" CASCADE;
DROP TABLE IF EXISTS "users" CASCADE;

DROP TYPE IF EXISTS role_enum CASCADE;
DROP TYPE IF EXISTS tingkat_kepentingan_enum CASCADE;
DROP TYPE IF EXISTS sumber_usulan_enum CASCADE;
DROP TYPE IF EXISTS status_usulan_enum CASCADE;
*/

-- 2. CREATE ENUM TYPES (dari fix_db.go & nuclear_fix.go)
CREATE TYPE role_enum AS ENUM ('pep', 'pptk', 'ppkom', 'pp', 'unit');
CREATE TYPE tingkat_kepentingan_enum AS ENUM ('Sangat Penting', 'Penting', 'Biasa');
CREATE TYPE sumber_usulan_enum AS ENUM ('Unit Kerja', 'Publik');
CREATE TYPE status_usulan_enum AS ENUM (
    'MENUNGGU_PEP', 
    'DISETUJUI_PEP', 
    'DIDISPOSISI_PPTK', 
    'DITERIMA_PPTK', 
    'DITERIMA_PPKOM', 
    'SETUJU_PPKOM', 
    'TOLAK_PPKOM', 
    'PROSES_PP', 
    'REALISASI_SELESAI', 
    'DIKEMBALIKAN_KE_PEP', 
    'DIKEMBALIKAN_KE_PPTK', 
    'DIDISPOSISI_PPKOM', 
    'DIDISPOSISI_PP', 
    'DIPERBAIKI_UNIT', 
    'GESER_TAHUN_DEPAN'
);

-- 3. ENSURE FALLBACK KATEGORI (dari fix_db.go & ultimate_fix_db.go)
INSERT INTO kategori_belanja (id, kode_kategori, nama_kategori, created_at) 
VALUES (1, 'K-001', 'Belanja Barang dan Jasa', NOW()) 
ON CONFLICT (id) DO NOTHING;

-- 4. ALTER TABLES (dari ultimate_fix_db.go)
ALTER TABLE usulan ADD COLUMN IF NOT EXISTS foto_barang text;
ALTER TABLE usulan ALTER COLUMN foto_barang TYPE varchar(255);
