--
-- PostgreSQL database dump
--

-- Dumped from database version 18.3
-- Dumped by pg_dump version 18.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

-- ENUM TYPES
CREATE TYPE public.pergeseran_enum AS ENUM (
    'TAMBAH',
    'KURANG'
);

CREATE TYPE public.role_enum AS ENUM (
    'pep',
    'pptk',
    'ppkom',
    'pp',
    'unit'
);

CREATE TYPE public.status_usulan_enum AS ENUM (
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

CREATE TYPE public.sumber_usulan_enum AS ENUM (
    'Unit Kerja',
    'Publik'
);

CREATE TYPE public.tingkat_kepentingan_enum AS ENUM (
    'Sangat Penting',
    'Penting',
    'Biasa'
);

-- TABLES
CREATE TABLE public.kategori_belanja (
    id bigserial PRIMARY KEY,
    kode_kategori character varying(50) NOT NULL UNIQUE,
    nama_kategori character varying(150) NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE public.users (
    id bigserial PRIMARY KEY,
    nama character varying(150) NOT NULL,
    username character varying(50) NOT NULL UNIQUE,
    password_hash character varying(255) NOT NULL,
    role public.role_enum NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE public.detail_anggaran (
    id bigserial PRIMARY KEY,
    category_id bigint NOT NULL REFERENCES public.kategori_belanja(id),
    nama character varying(255) NOT NULL,
    nominal numeric(15,2) NOT NULL,
    pptk_id bigint NOT NULL REFERENCES public.users(id),
    ppkom_id bigint REFERENCES public.users(id),
    pp_id bigint REFERENCES public.users(id),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE public.master_rincian_belanja (
    id bigserial PRIMARY KEY,
    kategori_belanja_id bigint NOT NULL REFERENCES public.kategori_belanja(id),
    nama_barang character varying(255) NOT NULL,
    spesifikasi text,
    satuan character varying(50),
    harga_referensi numeric(15,2) NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE public.usulan (
    id bigserial PRIMARY KEY,
    kode_tiket character varying(50) NOT NULL UNIQUE,
    nama_pengusul character varying(150) NOT NULL,
    no_hp character varying(20) NOT NULL,
    unit_ruangan character varying(100),
    nama_usulan character varying(200) NOT NULL,
    keterangan text,
    jumlah bigint NOT NULL,
    satuan character varying(50),
    kegentingan character varying(50) NOT NULL,
    tingkat_kepentingan public.tingkat_kepentingan_enum NOT NULL,
    sumber_usulan public.sumber_usulan_enum DEFAULT 'Unit Kerja'::public.sumber_usulan_enum NOT NULL,
    foto_barang character varying(255),
    kategori_belanja_id bigint NOT NULL REFERENCES public.kategori_belanja(id),
    status_kode public.status_usulan_enum DEFAULT 'MENUNGGU_PEP'::public.status_usulan_enum NOT NULL,
    pptk_user_id bigint REFERENCES public.users(id),
    ppkom_user_id bigint REFERENCES public.users(id),
    pp_user_id bigint REFERENCES public.users(id),
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    catatan_pep text,
    disposisi_pep_at timestamp with time zone,
    catatan_pptk text,
    disposisi_pptk_at timestamp with time zone,
    alasan_return text,
    return_at timestamp with time zone,
    catatan_ppkom text,
    setuju_ppkom_at timestamp with time zone,
    alasan_tolak text,
    reject_at timestamp with time zone,
    is_return_diketahui boolean DEFAULT false
);

CREATE TABLE public.riwayat_usulan (
    id bigserial PRIMARY KEY,
    usulan_id bigint NOT NULL REFERENCES public.usulan(id),
    actor_id bigint REFERENCES public.users(id),
    status_awal public.status_usulan_enum,
    status_akhir public.status_usulan_enum NOT NULL,
    catatan_alasan text,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE public.kewenangan_pegawai (
    id bigserial PRIMARY KEY,
    user_id bigint NOT NULL REFERENCES public.users(id),
    kategori_belanja_id bigint NOT NULL REFERENCES public.kategori_belanja(id),
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE public.realisasi_laporan (
    id bigserial PRIMARY KEY,
    usulan_id bigint NOT NULL UNIQUE REFERENCES public.usulan(id),
    pp_id bigint NOT NULL REFERENCES public.users(id),
    nama_vendor character varying(200) NOT NULL,
    nomor_kontrak character varying(100),
    harga_final numeric(15,2) NOT NULL,
    dokumen_laporan character varying(255),
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE public.usulan_detail_item (
    id serial PRIMARY KEY,
    usulan_id integer NOT NULL,
    master_barang_id integer,
    qty integer NOT NULL,
    harga_satuan numeric(15,2) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE public.histori_pergeseran_anggaran (
    id serial PRIMARY KEY,
    anggaran_id integer NOT NULL,
    actor_id integer NOT NULL,
    jenis_pergeseran public.pergeseran_enum NOT NULL,
    nominal numeric(15,2) NOT NULL,
    keterangan_alasan text NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);

-- INITIAL DATA SAMPLES (Transcribed from dump)
-- Inserting main roles
INSERT INTO public.users (id, nama, username, password_hash, role) VALUES
(1, 'Super Admin PEP', 'admin_pep', 'adminpep', 'pep'),
(2, 'Bekti Nugroho, S.Kom., M.Eng.', 'bekti', 'ppkom123', 'ppkom'),
(3, 'Suramin, S.KM., S.Kep., Ns.', 'suramin', 'ppkom123', 'ppkom'),
(4, 'dr. Nofi Kusumaningrum', 'nofi', 'ppkom123', 'ppkom'),
(5, 'Lilik Subagyo, S.Kep., Ns.', 'lilik', 'ppkom123', 'ppkom'),
(6, 'dr. Mayasari Ayu Hendrawati', 'mayasari', 'ppkom123', 'ppkom'),
(7, 'Eko Sujianto, S.Kep., Ns.', 'eko', 'ppkom123', 'ppkom'),
(8, 'Dr. dr. Kinik Darsono, M.Pd.Ked.', 'kinik', 'ppkom123', 'ppkom'),
(9, 'Imas Wulandari, S.Kom., M.Eng.', 'imas', 'ppkom123', 'ppkom'),
(10, 'Yuwono Laksito, A.Md.', 'yuwono', 'pptk123', 'pptk'),
(11, 'Andreas Hasta Riawan', 'andreas', 'pptk123', 'pptk'),
(12, 'Nunung Agus Dwi H., M.Si.', 'nunung', 'pptk123', 'pptk'),
(13, 'Fitria Purnamawati, S.Kep., Ns., M.Kep.', 'fitria', 'pp123', 'pp'),
(14, 'Agus Dwiyanto, S.Farm., Apt.', 'agus', 'pp123', 'pp'),
(15, 'Purwanto, A.Md.', 'purwanto', 'pp123', 'pp'),
(16, 'Bagus Prasetya Utama, A.Md.TE.', 'bagus', 'pp123', 'pp'),
(17, 'Candra Nugraheni, S.ST.', 'candra-nugraheni', 'pp123', 'pp'),
(18, 'Candra Rosanti, S.KM.', 'candra-rosanti', 'pp123', 'pp'),
(19, 'Tulus Wahyuno, A.Md.', 'tulus', 'pp123', 'pp')
ON CONFLICT (id) DO NOTHING;

-- Inserting Categories
INSERT INTO public.kategori_belanja (id, kode_kategori, nama_kategori) VALUES
(1, 'BBJ', 'Belanja Barang dan Jasa'),
(27, 'BMPM', 'Belanja Modal Peralatan dan Mesin'),
(28, 'BMGB', 'Belanja Modal Gedung dan Bangunan'),
(29, 'BMJI', 'Belanja Modal Jalan, Jaringan, dan Irigasi')
ON CONFLICT (id) DO NOTHING;

-- Synchronizing sequences (if needed manually)
SELECT pg_catalog.setval('public.kategori_belanja_id_seq', 29, true);
SELECT pg_catalog.setval('public.users_id_seq', 19, true);
