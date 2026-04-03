# Panduan Database SIUL PBJ

Folder ini berisi skema database dan data awal untuk memudahkan kolaborasi tim.

## Isi File
- `siul_pbj_full_dump.sql`: Backup lengkap database (Skema + Data).
- `schema_full.sql`: Skema database yang lebih bersih dengan seed data dasar.
- `import.bat`: Script sekali klik untuk pengguna Windows (pastikan PostgreSQL & `psql.exe` sudah terpasang dan ada di PATH).

## Cara Setup untuk Kolaborator Baru
1. Pastikan Anda sudah menginstal PostgreSQL.
2. Buat database baru bernama `siul_db` (atau sesuai DB_NAME di `.env`).
   ```sql
   CREATE DATABASE siul_db;
   ```
3. Jalankan perintah import di terminal dari folder ini:
   ```bash
   psql -h localhost -p 5432 -U postgres -d siul_db -f siul_pbj_full_dump.sql
   ```
   *Atau klik dua kali pada `import.bat` jika menggunakan Windows.*

4. Sesuaikan file `.env` di root folder backend dengan kredensial PostgreSQL lokal Anda.
