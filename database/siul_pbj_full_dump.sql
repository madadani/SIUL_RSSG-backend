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

--
-- TOC entry 873 (class 1247 OID 17384)
-- Name: pergeseran_enum; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.pergeseran_enum AS ENUM (
    'TAMBAH',
    'KURANG'
);


ALTER TYPE public.pergeseran_enum OWNER TO postgres;

--
-- TOC entry 882 (class 1247 OID 18338)
-- Name: role_enum; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.role_enum AS ENUM (
    'pep',
    'pptk',
    'ppkom',
    'pp',
    'unit'
);


ALTER TYPE public.role_enum OWNER TO postgres;

--
-- TOC entry 891 (class 1247 OID 18364)
-- Name: status_usulan_enum; Type: TYPE; Schema: public; Owner: postgres
--

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


ALTER TYPE public.status_usulan_enum OWNER TO postgres;

--
-- TOC entry 888 (class 1247 OID 18358)
-- Name: sumber_usulan_enum; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.sumber_usulan_enum AS ENUM (
    'Unit Kerja',
    'Publik'
);


ALTER TYPE public.sumber_usulan_enum OWNER TO postgres;

--
-- TOC entry 885 (class 1247 OID 18350)
-- Name: tingkat_kepentingan_enum; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.tingkat_kepentingan_enum AS ENUM (
    'Sangat Penting',
    'Penting',
    'Biasa'
);


ALTER TYPE public.tingkat_kepentingan_enum OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 237 (class 1259 OID 18569)
-- Name: detail_anggaran; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.detail_anggaran (
    id bigint NOT NULL,
    category_id bigint NOT NULL,
    nama character varying(255) NOT NULL,
    nominal numeric(15,2) NOT NULL,
    pptk_id bigint NOT NULL,
    ppkom_id bigint,
    pp_id bigint,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.detail_anggaran OWNER TO postgres;

--
-- TOC entry 220 (class 1259 OID 17499)
-- Name: histori_pergeseran_anggaran; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.histori_pergeseran_anggaran (
    id integer NOT NULL,
    anggaran_id integer NOT NULL,
    actor_id integer NOT NULL,
    jenis_pergeseran public.pergeseran_enum NOT NULL,
    nominal numeric(15,2) NOT NULL,
    keterangan_alasan text NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.histori_pergeseran_anggaran OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 17498)
-- Name: histori_pergeseran_anggaran_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.histori_pergeseran_anggaran_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.histori_pergeseran_anggaran_id_seq OWNER TO postgres;

--
-- TOC entry 5040 (class 0 OID 0)
-- Dependencies: 219
-- Name: histori_pergeseran_anggaran_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.histori_pergeseran_anggaran_id_seq OWNED BY public.histori_pergeseran_anggaran.id;


--
-- TOC entry 226 (class 1259 OID 18410)
-- Name: kategori_belanja; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.kategori_belanja (
    id bigint NOT NULL,
    kode_kategori character varying(50) NOT NULL,
    nama_kategori character varying(150) NOT NULL,
    created_at timestamp with time zone
);


ALTER TABLE public.kategori_belanja OWNER TO postgres;

--
-- TOC entry 225 (class 1259 OID 18409)
-- Name: kategori_belanja_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.kategori_belanja_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.kategori_belanja_id_seq OWNER TO postgres;

--
-- TOC entry 5041 (class 0 OID 0)
-- Dependencies: 225
-- Name: kategori_belanja_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.kategori_belanja_id_seq OWNED BY public.kategori_belanja.id;


--
-- TOC entry 234 (class 1259 OID 18503)
-- Name: kewenangan_pegawai; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.kewenangan_pegawai (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    kategori_belanja_id bigint NOT NULL,
    created_at timestamp with time zone
);


ALTER TABLE public.kewenangan_pegawai OWNER TO postgres;

--
-- TOC entry 233 (class 1259 OID 18502)
-- Name: kewenangan_pegawai_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.kewenangan_pegawai_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.kewenangan_pegawai_id_seq OWNER TO postgres;

--
-- TOC entry 5042 (class 0 OID 0)
-- Dependencies: 233
-- Name: kewenangan_pegawai_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.kewenangan_pegawai_id_seq OWNED BY public.kewenangan_pegawai.id;


--
-- TOC entry 228 (class 1259 OID 18422)
-- Name: master_rincian_belanja; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.master_rincian_belanja (
    id bigint NOT NULL,
    kategori_belanja_id bigint NOT NULL,
    nama_barang character varying(255) NOT NULL,
    spesifikasi text,
    satuan character varying(50),
    harga_referensi numeric(15,2) NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.master_rincian_belanja OWNER TO postgres;

--
-- TOC entry 227 (class 1259 OID 18421)
-- Name: master_rincian_belanja_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.master_rincian_belanja_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.master_rincian_belanja_id_seq OWNER TO postgres;

--
-- TOC entry 5043 (class 0 OID 0)
-- Dependencies: 227
-- Name: master_rincian_belanja_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.master_rincian_belanja_id_seq OWNED BY public.master_rincian_belanja.id;


--
-- TOC entry 236 (class 1259 OID 18523)
-- Name: realisasi_laporan; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.realisasi_laporan (
    id bigint NOT NULL,
    usulan_id bigint NOT NULL,
    pp_id bigint NOT NULL,
    nama_vendor character varying(200) NOT NULL,
    nomor_kontrak character varying(100),
    harga_final numeric(15,2) NOT NULL,
    dokumen_laporan character varying(255),
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.realisasi_laporan OWNER TO postgres;

--
-- TOC entry 235 (class 1259 OID 18522)
-- Name: realisasi_laporan_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.realisasi_laporan_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.realisasi_laporan_id_seq OWNER TO postgres;

--
-- TOC entry 5044 (class 0 OID 0)
-- Dependencies: 235
-- Name: realisasi_laporan_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.realisasi_laporan_id_seq OWNED BY public.realisasi_laporan.id;


--
-- TOC entry 232 (class 1259 OID 18491)
-- Name: riwayat_usulan; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.riwayat_usulan (
    id bigint NOT NULL,
    usulan_id bigint NOT NULL,
    actor_id bigint,
    status_awal public.status_usulan_enum,
    status_akhir public.status_usulan_enum NOT NULL,
    catatan_alasan text,
    created_at timestamp with time zone
);


ALTER TABLE public.riwayat_usulan OWNER TO postgres;

--
-- TOC entry 231 (class 1259 OID 18490)
-- Name: riwayat_usulan_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.riwayat_usulan_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.riwayat_usulan_id_seq OWNER TO postgres;

--
-- TOC entry 5045 (class 0 OID 0)
-- Dependencies: 231
-- Name: riwayat_usulan_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.riwayat_usulan_id_seq OWNED BY public.riwayat_usulan.id;


--
-- TOC entry 224 (class 1259 OID 18396)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    nama character varying(150) NOT NULL,
    username character varying(50) NOT NULL,
    password_hash character varying(255) NOT NULL,
    role public.role_enum NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 18395)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- TOC entry 5046 (class 0 OID 0)
-- Dependencies: 223
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 230 (class 1259 OID 18457)
-- Name: usulan; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.usulan (
    id bigint NOT NULL,
    kode_tiket character varying(50) NOT NULL,
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
    kategori_belanja_id bigint NOT NULL,
    status_kode public.status_usulan_enum DEFAULT 'MENUNGGU_PEP'::public.status_usulan_enum NOT NULL,
    pptk_user_id bigint,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
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
    ppkom_user_id bigint,
    pp_user_id bigint,
    is_return_diketahui boolean DEFAULT false
);


ALTER TABLE public.usulan OWNER TO postgres;

--
-- TOC entry 222 (class 1259 OID 17551)
-- Name: usulan_detail_item; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.usulan_detail_item (
    id integer NOT NULL,
    usulan_id integer NOT NULL,
    master_barang_id integer,
    qty integer NOT NULL,
    harga_satuan numeric(15,2) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.usulan_detail_item OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 17550)
-- Name: usulan_detail_item_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.usulan_detail_item_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.usulan_detail_item_id_seq OWNER TO postgres;

--
-- TOC entry 5047 (class 0 OID 0)
-- Dependencies: 221
-- Name: usulan_detail_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.usulan_detail_item_id_seq OWNED BY public.usulan_detail_item.id;


--
-- TOC entry 229 (class 1259 OID 18456)
-- Name: usulan_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.usulan_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.usulan_id_seq OWNER TO postgres;

--
-- TOC entry 5048 (class 0 OID 0)
-- Dependencies: 229
-- Name: usulan_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.usulan_id_seq OWNED BY public.usulan.id;


--
-- TOC entry 4814 (class 2604 OID 17502)
-- Name: histori_pergeseran_anggaran id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.histori_pergeseran_anggaran ALTER COLUMN id SET DEFAULT nextval('public.histori_pergeseran_anggaran_id_seq'::regclass);


--
-- TOC entry 4819 (class 2604 OID 18413)
-- Name: kategori_belanja id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.kategori_belanja ALTER COLUMN id SET DEFAULT nextval('public.kategori_belanja_id_seq'::regclass);


--
-- TOC entry 4826 (class 2604 OID 18506)
-- Name: kewenangan_pegawai id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.kewenangan_pegawai ALTER COLUMN id SET DEFAULT nextval('public.kewenangan_pegawai_id_seq'::regclass);


--
-- TOC entry 4820 (class 2604 OID 18425)
-- Name: master_rincian_belanja id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.master_rincian_belanja ALTER COLUMN id SET DEFAULT nextval('public.master_rincian_belanja_id_seq'::regclass);


--
-- TOC entry 4827 (class 2604 OID 18526)
-- Name: realisasi_laporan id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.realisasi_laporan ALTER COLUMN id SET DEFAULT nextval('public.realisasi_laporan_id_seq'::regclass);


--
-- TOC entry 4825 (class 2604 OID 18494)
-- Name: riwayat_usulan id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.riwayat_usulan ALTER COLUMN id SET DEFAULT nextval('public.riwayat_usulan_id_seq'::regclass);


--
-- TOC entry 4818 (class 2604 OID 18399)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 4821 (class 2604 OID 18460)
-- Name: usulan id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.usulan ALTER COLUMN id SET DEFAULT nextval('public.usulan_id_seq'::regclass);


--
-- TOC entry 4816 (class 2604 OID 17554)
-- Name: usulan_detail_item id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.usulan_detail_item ALTER COLUMN id SET DEFAULT nextval('public.usulan_detail_item_id_seq'::regclass);

-- ... (Rest of COPY the user provided) ...
